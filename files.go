package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const data_dir = "data"

func writeString(filename, content string) {
	fmt.Println("Writing file")
	path := filepath.Join(data_dir, filename)
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	length, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %s", file.Name())
	fmt.Printf("\nfile length: %d\n", length)
}

func readString(filename string) (string, error) {
	fmt.Println("Reading file")
	path := filepath.Join(data_dir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	fmt.Println("file name " + filename)
	fmt.Printf("file size %d\n", len(data))
	fmt.Printf("file content : %s\n", data)
	return string(data), nil
}
