package main

import (
	"os"
	"path/filepath"
)

type PageManager struct {
	file *os.File
}

func PageManagerCreator(dbPath string) (*PageManager, error) {
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(dbPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	pm := &PageManager{file: file}

	return pm, nil
}
