package main

import (
	"fmt"
	"Masker/utils"
	"os"
)

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
	prod := &utils.FileProducer{FilePath: input}
	pres := &utils.FilePresenter{Filepath: output}

	service := &utils.Service{Prod: prod, Pres: pres}

	service.Run()
}
