package main

import (
	"testing"
)

const (
	DBFilename = "app.db"
)

func TestPageManagerCreator(t *testing.T) {
	pm, err := PageManagerCreator(DBFilename)
	if err != nil {
		panic(err)
	}
	defer pm.file.Close()

	fileInfo, err := pm.file.Stat()
	if err != nil {
		t.Fatalf("failed to stat db file: %v", err)
	}
	if fileInfo.IsDir() {
		t.Fatalf("%s exists but is a directory, expected a file", DBFilename)
	}
	if fileInfo.Size() == 0 {
		t.Fatalf("%s was created but is empty", DBFilename)
	}
}
