package main

import (
	"fmt"
	"time"
)

func main() {
	var err error

	fmt.Println("Starting GoRoutine!!!")
	go sayHello()

	go printNumbers()
	go printLetters()

	go func() {
		err = doWork()
	}()

	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Work completed successfully!!!")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Microsecond)
	}
}

func printLetters() {
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Println(string(letter))
		time.Sleep(100 * time.Microsecond)
	}
}

func doWork() error {
	//simulate a work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("Error occurred in do work!!!")
}
