package utils

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present([]string) error
}

type Service struct {
	Prod Producer
	Pres Presenter
}

func (r *Service) Run() {
	data, err := r.Prod.Produce()
	if err != nil {
		fmt.Printf("Error occured while producing: %v\n", err)
		return
	}

	ch := make(chan string)
	var wg sync.WaitGroup
	for _, fData := range data {
		wg.Add(1)
		go Find(fData, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var chArray []string
	for res := range ch {
		chArray = append(chArray, res)
	}
	if err := r.Pres.Present(chArray); err != nil {
		fmt.Printf("Error occured in presenter: %v\n", err)
		return
	}
}

type FileProducer struct {
	FilePath string
}

func (f *FileProducer) Produce() ([]string, error) {
	data, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read a file: %v", err)
	}
	str := strings.Split(string(data), "\n")
	return str, nil
}

type FilePresenter struct {
	Filepath string
}

func (f *FilePresenter) Present(data []string) error {
	file, err := os.Create(f.Filepath)
	if err != nil {
		return fmt.Errorf("failed to create a file: %v", err)
	}
	defer file.Close()
	for _, r := range data {
		if _, err := file.WriteString(r); err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}
	return nil
}
