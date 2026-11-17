package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createFile(destination string, fileName string) (*os.File, error) {
	if err := os.MkdirAll(destination, 0755); err != nil {
		return nil, fmt.Errorf("create output folder: %w", err)
	}

	destinationPath := filepath.Join(destination, fileName)
	file, err := os.Create(destinationPath)
	if err != nil {
		return nil, fmt.Errorf("create output file: %w", err)
	}

	return file, nil
}
