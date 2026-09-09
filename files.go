package main

import (
	"os"
	"path/filepath"
)

const data_dir = "data"

func writeString(filename, content string) (int, error) {
	path := filepath.Join(data_dir, filename)
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return file.WriteString(content)
}

func readString(filename string) (string, error) {
	path := filepath.Join(data_dir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}



func init() {
	if err := os.MkdirAll(data_dir, 0755); err != nil {
        panic("We could not create the data directory")
    }
}
