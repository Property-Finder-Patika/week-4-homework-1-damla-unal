package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*

Exercise:
	How would you implement a program to list go source code files in a given directory including all sub-directories recursively?
	Please use filepath.WalkDir() function (https://pkg.go.dev/path/filepath#WalkDir) for this purpose.

*/

func main() {

	//The path/filepath package provides a handy way to scan all the files in a directory,
	//it will automatically scan each sub-directories in the directory.
	err := filepath.Walk("/Users/damlaunal/Projects/Patika/week-4-homework-1-damla-unal",
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		return
	}

}
