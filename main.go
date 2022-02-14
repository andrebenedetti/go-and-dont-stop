package main

import (
	"fmt"
	"time"
)

func watch() {
	oldFiles := RecursivelyReadDirectoryGoFiles()
	for {
		files := RecursivelyReadDirectoryGoFiles()
		if FilesChanged(oldFiles, files) {
			fmt.Println("files changed")
		}
		oldFiles = files
		time.Sleep(time.Second * 2)
	}
}


func main() {
	watch()
}