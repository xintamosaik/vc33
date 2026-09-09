package main

import (
	"os"
	"path/filepath"
)

const dataDir = "data"

func writeString(filename, content string) error {
	path := filepath.Join(dataDir, filename)
	return os.WriteFile(path, []byte(content), 0644)
}

func readString(filename string) (string, error) {
	path := filepath.Join(dataDir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func init() {
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		panic("We could not create the data directory")
	}
}
