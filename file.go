package main

import (
	"hash/fnv"
	"io/fs"
	"os"
	"strings"
)

type FileMetadata struct {
	name  string
	sum64 uint64
}

type FileHashMap map[string]FileMetadata

func HasNumberOfFilesChanged(oldFiles, newFiles FileHashMap) bool {
	return len(oldFiles) != len(newFiles)
}

// renamed?
func HasFileContentChanged(oldFiles FileHashMap, file FileMetadata) bool {
	return oldFiles[file.name].sum64 == file.sum64
}

func FilesChanged(oldFiles, newFiles FileHashMap) bool {
	HasFileContentChanged(oldFiles, newFiles["main.go"])
	return HasNumberOfFilesChanged(oldFiles, newFiles)
}

func RecursivelyReadDirectoryGoFiles() FileHashMap {
	dirFs := os.DirFS(".")
	dir, _ := fs.ReadDir(dirFs, ".")

	var files = make(map[string]FileMetadata)

	for _, val := range dir {
		if strings.HasSuffix(val.Name(), ".go") {
			hash := fnv.New64()
			data, err := os.ReadFile(val.Name())
			hash.Write(data)
			if err != nil {
				panic(err)
			}
			x := FileMetadata{val.Name(), hash.Sum64()}
			files[val.Name()] = x
		}
	}
	return files
}
