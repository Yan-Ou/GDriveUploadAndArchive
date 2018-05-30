package main

import (
	"fmt"
	//"io/ioutil"
	//"log"
	"os"
	"path/filepath"
	"sort"
)

type FileData struct {
	File os.FileInfo
	Path string
}

func prepareFiles() []FileData {
	dir := "/Users/yan/Workspace/dgt/photos"
	files := []FileData{}

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}

		fileData := FileData{
			File: f,
			Path: path,
		}

		if filepath.Ext(fileData.Path) == ".jpg" {
			files = append(files, fileData)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].File.ModTime().After(files[j].File.ModTime())
	})

	for _, fis := range files {
		fmt.Printf("File name: %s, Last modified time: %s\n", fis.File.Name(), fis.File.ModTime())
	}

	return files
}
