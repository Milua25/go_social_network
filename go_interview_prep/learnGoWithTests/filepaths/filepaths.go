package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	absPath, err := filepath.Abs("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Absolute path:", absPath)

	//join paths
	joinedPath := filepath.Join("test", "test.txt")
	fmt.Println("Joined path:", joinedPath)

	//split paths
	dir, file := filepath.Split(absPath)
	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)

	//basename
	basename := filepath.Base(absPath)
	fmt.Println("Basename:", basename)
}
