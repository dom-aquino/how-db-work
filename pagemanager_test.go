package main

import (
	"testing"
)

const (
	DBFilename = "app.db"
)

func TestPageManager(t *testing.T) {
	pm, err := PageManagerCreator(DBFilename)
	if err != nil {
		panic(err)
	}
	defer pm.file.Close()
}
