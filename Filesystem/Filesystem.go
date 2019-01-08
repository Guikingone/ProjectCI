package Filesystem

import (
	"log"
	"os"
)

type Filesystem struct {
	lastOperation string
}

func (filesystem *Filesystem) CreateDir(filePath string) error {

	_, fsErr := os.Stat(filePath)

	if os.IsExist(fsErr) {
		return fsErr
	}

	err := os.MkdirAll(filePath, 0755)


	if err != nil {
		log.Printf("Creating %s", filePath)
	}

	return err
}
