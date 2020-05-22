package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// FileHash type definition
type FileHash struct {
	path         string
	allFiles     []string
	fileNamesMap map[string][]string
}

// FindAllFiles gathers all the files in a nested directory
func (fh *FileHash) FindAllFiles() error {
	err := filepath.Walk(fh.path, func(path string, f os.FileInfo, err error) error {
		fh.allFiles = append(fh.allFiles, path)
		return nil
	})

	return err
}

// BuildFileNamesMap a map of hash of contents and a list of file names that hash to the same value
func (fh *FileHash) BuildFileNamesMap() error {
	for _, fn := range fh.allFiles {
		hash := md5.New()
		fi, err := os.Stat(fn)
		if err == nil {
			if fi.Mode().IsRegular() {
				dat, err := ioutil.ReadFile(fn)
				if err == nil {
					if len(fh.fileNamesMap[string(hash.Sum([]byte(dat)))]) == 0 {
						fh.fileNamesMap[string(hash.Sum([]byte(dat)))] = []string{fn}
					} else {
						fh.fileNamesMap[string(hash.Sum([]byte(dat)))] = append(fh.fileNamesMap[string(hash.Sum([]byte(dat)))], fn)
					}
				} else {
					return err
				}
			}
		} else {
			return err
		}
	}
	return nil
}

// GetIdenticalFiles gets files are mapped to the same hash
func (fh FileHash) GetIdenticalFiles() {
	for _, v := range fh.fileNamesMap {
		if len(v) > 1 {
			fmt.Println("Identical files:", v)
		}
	}
}

// New instance of FileHash
func New(path string) *FileHash {
	return &FileHash{
		path:         path,
		fileNamesMap: make(map[string][]string),
	}
}

func main() {

	// Driver code
	fileHash := New("test")

	err := fileHash.FindAllFiles()

	if err != nil {
		fmt.Print("Error while walking the directory", err)
	}

	err = fileHash.BuildFileNamesMap()

	if err != nil {
		fmt.Printf("Error while building the filenames map", err)
	}

	fileHash.GetIdenticalFiles()

	// >>> go run dupe_files.go                                                                                                                                   (*master+6) 09:23:46
	// Identical files: [test/a test/b.txt test/test1/c test/test1/test2/e test/test1/w test/test1/x test/test1/y test/test1/z]
	// Identical files: [test/c.txt test/d.txt test/test1/test2/f]
}
