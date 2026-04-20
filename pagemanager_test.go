package main

import (
	"testing"
)

const (
	DBFilename = "app.db"
)

func TestPageManagerCreator(t *testing.T) {
	// Test #1: Check the creation of app.db file
	pm, err := PageManagerCreator(DBFilename)
	if err != nil {
		panic(err)
	}
	defer pm.file.Close()

	// Test #2: Check the allocation of memory for the new page
	newPageID, err := pm.AllocatePage()
	if err != nil {
		t.Fatal(err.Error())
	}

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

	// Test #3: Check that the page was created
	var newPage Page
	copy(newPage[:], []byte("Page one data..."))

	err = pm.WritePage(newPageID, &newPage)
	if err != nil {
		t.Fatalf("Failed to write page: %s", err)
	}

	var readPage Page
	err = pm.ReadPage(newPageID, &readPage)
	if err != nil {
		t.Fatalf("Failed to read page: %s", err)
	}

	anotherPageID, err := pm.AllocatePage()
	if err != nil {
		t.Fatal(err.Error())
	}

	var anotherNewPage Page
	copy(anotherNewPage[:], []byte("Page two data..."))

	err = pm.WritePage(anotherPageID, &anotherNewPage)
	if err != nil {
		t.Fatalf("Failed to write page: %s", err)
	}

	var anotherReadPage Page
	err = pm.ReadPage(anotherPageID, &anotherReadPage)
	if err != nil {
		t.Fatalf("Failed to read page: %s", err)
	}

	/*  How to interpret the hex data (hexdump -C app.db | head)
	| Byte offset | 32 bytes payload | ASCII printable characters |
	00000000  00 00 00 01 00 00 00 00  00 00 00 00 00 00 00 00  |................|
	00000010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
	*
	00001000  50 61 67 65 20 6f 6e 65  20 64 61 74 61 2e 2e 2e  |Page one data...|
	00001010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
	*
	00002000
	*/
}
