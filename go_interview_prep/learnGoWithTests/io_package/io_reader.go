package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// function to demonstrate the use of io package to read data
func readFromReader(r io.Reader) ([]byte, error) {
	// buf := new(bytes.Buffer)
	buf := make([]byte, 1024) // create a buffer to hold the data
	n, err := r.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error reading from the reader:", err)
	}
}

// close Resource
func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // creates memory on the stack
	buf.WriteString("Hello Buffer!")

	fmt.Println(buf.String())
}

// multiReader

func multiReader() {
	r1 := strings.NewReader("House ")
	r2 := strings.NewReader("Of ")
	r3 := strings.NewReader("Dragon!!!!!")

	mr := io.MultiReader(r1, r2, r3)

	buf := new(bytes.Buffer) // allocates memory on the heap

	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln(err)
	}

	line := buf.String()
	fmt.Println(line)

}

func readFromPipe() {
	rr, rw := io.Pipe()

	errCh := make(chan error, 1)

	go func() {
		defer closeResource(rw)
		defer close(errCh)
		_, err := rw.Write([]byte("Hello Piper!!!"))
		if err != nil {
			errCh <- err
		}

	}()

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(rr)
	if err != nil {
		log.Fatalln(err)
	}

	if err := <-errCh; err != nil {
		log.Fatalf("go routine failed: %v", err)
	}

	fmt.Println(buf.String())

}

func writeToFile(filepath, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file:", err)
	}
	defer closeResource(file)

	// _, err = file.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// to convert Type(value)

	writer := io.Writer(file)
	_, err = writer.Write([]byte(data))
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	fmt.Println("IO Package")

	// Example of using readFromReader with a string reader
	fmt.Println("====New Reader====")
	data := "Hello, World!"
	newReader := strings.NewReader(data)
	result, err := readFromReader(newReader)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(result))

	fmt.Println("====Multi Reader====")
	multiReader()

	fmt.Println("====Write to Writer====")
	var writer bytes.Buffer
	writeToWriter(&writer, data)
	fmt.Println(writer.String())

	fmt.Println("====Write to Writer====")
	bufferExample()

	fmt.Println("====Pipe Example====")
	readFromPipe()

	fmt.Println("====Write to File====")
	writeToFile("sample.txt", data+"\n")

}
