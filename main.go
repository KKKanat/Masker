package main

import (
	"fmt"
	"os"
	// "path/filepath"
)

type producer interface {
	produce() (string, error)
}

type presenter interface {
	present(string) error
}

type Service struct{
	prod producer
	pres presenter
}

type FileProducer struct {
	filePath string
}

func (f *FileProducer) produce() (string, error) {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return "", fmt.Errorf("Failed to read a file: %v", err)
	}
	str := string(data)
	return str, nil
}

type FilePresenter struct {
	filepath string
}

func (f *FilePresenter) present(data string) error {
	file, err := os.Create(f.filepath)
	if err != nil {
		return fmt.Errorf("Failed to create a file: %v", err)
	}
	defer file.Close()

	// for _, r := range data {
		if _, err := file.WriteString(data); err != nil {
			return fmt.Errorf("Failed to write to file: %v", err)
		}
	// }
	return nil
}

func runeLength(args string) int {
	count := 0
	for range args {
		count++
	}
	return count
}

func find(text string) string {
	pattern := "http://"
	patternLength := runeLength(pattern)
	needToMask := false
	textInRunes := []rune(text)
	for i, r := range textInRunes {
		if r == ' ' {
			needToMask = false
		}
		if needToMask {
			textInRunes[i] = '*'
		}
		if !needToMask && i+1 >= patternLength {
			if pattern == text[i+1-patternLength:i+1] {
				needToMask = true
			}
		}
	}
	return string(textInRunes)
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Didn't match the usage input: go run main.go inputFile outputFile[optional]")
		return
	}
	input := os.Args[1]
	output := "output.txt"
	if len(os.Args) == 3 {
		output = os.Args[2]
	}
	prod := &FileProducer{filePath: input}
	pres := &FilePresenter{filepath: output}
	
	service := &Service{prod: prod, pres: pres}

	data, err := service.prod.produce()
	if err != nil{
		fmt.Printf("Error occured while producing: %v\n", err)
		return
	}

	data = find(data)

	if err := service.pres.present(data); err != nil {
		fmt.Printf("Error occured in presenter: %v\n", err)
	}
}
