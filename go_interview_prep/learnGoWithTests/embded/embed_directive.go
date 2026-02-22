package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed test.txt
var embeddedFile string

//go:embed basics
var basicsFolder embed.FS

// check errors
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(embeddedFile)

	baseFiles, err := basicsFolder.ReadDir("basics")
	checkError(err)
	for _, file := range baseFiles {
		fmt.Println(file.Name())
	}

	content, err := basicsFolder.ReadFile("basics/hello.txt")
	checkError(err)

	fmt.Println(string(content))

	err = fs.WalkDir(basicsFolder, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Println(path, d.IsDir())
		return nil
	})
	checkError(err)

}
