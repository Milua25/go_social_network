package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greeting(writer io.Writer, name string) {
	_, err := fmt.Fprintf(writer, "Hello, %v!\n", name)
	if err != nil {
		return
	}
	//fmt.Println()
}

func MyGreetingHandler(w http.ResponseWriter, req *http.Request) {
	Greeting(w, "Elodie")
}

func main() {
	Greeting(os.Stdout, "Elodie")

	log.Fatalln(http.ListenAndServe(":5051", http.HandlerFunc(MyGreetingHandler)))
}
