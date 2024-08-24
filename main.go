package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func main() {
	d := "C:\\Users\\bisqu\\Documents\\Workspaces\\Personal"

	err := getExts(d)

	if err != nil {
		log.Fatal(err)
	}
}

func getExts(d string) error {

	err := filepath.Walk(d, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)

			if ext != "" {
				fmt.Printf("%s => %s", path, ext)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
