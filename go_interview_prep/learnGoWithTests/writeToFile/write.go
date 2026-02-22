package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// // write to the file
	// _, err = file.WriteString("Hello, World!")
	// if err != nil {
	// 	panic(err)
	// }

	// read from the file
	// content, err := os.ReadFile("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(content))

	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read from the file line by line
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("File created and written to")

	// line filtering
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Lorem") {
			fmt.Println(line)
		}

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
