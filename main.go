package main

import (
	"fmt"
	"hash/fnv"
	"io/fs"
	"os"
	"strings"
)



func main() {
	dirFs := os.DirFS(".")
	dir, _ := fs.ReadDir(dirFs, ".")

	for _,val := range dir {
		if strings.HasSuffix(val.Name(), ".go") {
			hash := fnv.New64()
			data, err := os.ReadFile(val.Name())
			hash.Write(data)
			if err != nil {
				panic(err)
			}
			x := FileMetadata{val.Name(), hash.Sum64()}
			fmt.Println(x)
		}
	}
}