package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
)

const (
	PageSize  = 4096
	MaxPageID = 1<<32 - 2
)

type Page [PageSize]byte

type PageManager struct {
	file       *os.File
	numOfPages uint32
}

func (pm *PageManager) AllocatePage() (uint32, error) {
	newPageID := pm.numOfPages + 1
	if newPageID+1 > MaxPageID {
		return newPageID, fmt.Errorf("Max number of pages is reached")
	}
	pm.numOfPages = newPageID
	offset := int64(newPageID) * PageSize

	if err := pm.file.Truncate(offset + PageSize); err != nil {
		return newPageID, err
	}

	header := make([]byte, PageSize)
	binary.BigEndian.PutUint32(header, newPageID)
	_, err := pm.file.WriteAt(header, 0)
	if err != nil {
		return newPageID, err
	}

	return newPageID, nil
}

// Writes "page" to disk
func (pm *PageManager) WritePage(pageID uint32, page *Page) error {
	offset := int64(pageID) * PageSize
	_, err := pm.file.WriteAt(page[:], offset)
	return err
}

// Loads page data to "page" for reading
func (pm *PageManager) ReadPage(pageID uint32, page *Page) error {
	if pageID == 0 || pageID > pm.numOfPages {
		return fmt.Errorf("Invalid pageID")
	}
	offset := int64(pageID) * PageSize
	_, err := pm.file.ReadAt(page[:], offset)
	return err
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
