package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileHash type definition
type FileHash struct {
	path         string
	allFiles     []string
	fileNamesMap map[string]string
}

// FindFiles gathers all the files in a nested directory
func (fh *FileHash) FindFiles() error {
	err := filepath.Walk(fh.path, func(path string, f os.FileInfo, err error) error {
		fh.allFiles = append(fh.allFiles, path)
		return nil
	})

	return err
}

// BuildFileNamesMap builds a map of the hash against a list of file names that hash to the same value
func (fh *FileHash) BuildFileNamesMap() {
	// todo
}

// GetIdenticalFiles gets files are mapped to the same hash
func (fh FileHash) GetIdenticalFiles() {
	// todo
}

func main() {

	// Driver code
	fileHash := FileHash{path: "test/test1"}

	err := fileHash.FindFiles()
	fileHash.BuildFileNamesMap()

	if err == nil {
		fmt.Print("All files: ", fileHash.allFiles)
	} else {
		fmt.Print("Error while walking the directory", err)
	}
}
