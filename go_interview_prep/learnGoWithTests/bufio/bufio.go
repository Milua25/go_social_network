package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// new reader

	reader := strings.NewReader("Hello, World!\n")

	_, err := io.Copy(os.Stdout, reader) //reader.WriteTo(os.Stdout)
	if err != nil {
		panic(err)
	}

	// create a buffer reader

	bufferedReader := bufio.NewReader(strings.NewReader("This is a buffered reader!"))
	// _, err = io.Copy(os.Stdout, bufferedReader)
	// if err != nil {
	// 	panic(err)
	// }

	data := make([]byte, 20)
	_, err = bufferedReader.Read(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// read string
	stringData, err := bufferedReader.ReadString('!') // read until the first occurrence of '!'
	if err != nil {
		panic(err)
	}
	fmt.Println(stringData)

	// write to a file
	writer := bufio.NewWriter(os.Stdout)

	_, err = writer.WriteString("This is a buffered writer!")
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

//
