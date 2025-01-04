package utils

import (
	"fmt"
	"os"
)

type Producer interface {
	Produce() (string, error)
}

type Presenter interface {
	Present(string) error
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

	data = Find(data)

	if err := r.Pres.Present(data); err != nil {
		fmt.Printf("Error occured in presenter: %v\n", err)
	}
}

type FileProducer struct {
	FilePath string
}

func (f *FileProducer) Produce() (string, error) {
	data, err := os.ReadFile(f.FilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read a file: %v", err)
	}
	str := string(data)
	return str, nil
}

type FilePresenter struct {
	Filepath string
}

func (f *FilePresenter) Present(data string) error {
	file, err := os.Create(f.Filepath)
	if err != nil {
		return fmt.Errorf("failed to create a file: %v", err)
	}
	defer file.Close()
	if _, err := file.WriteString(data); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}
