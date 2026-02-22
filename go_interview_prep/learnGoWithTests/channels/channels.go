package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// make(chan type)

// 	// greeting := make(chan []string) //unbuffered channels

// 	// greetString := []string{"Hello", "World", "Friend"}

// 	// go func() {
// 	// 	greeting <- greetString
// 	// }()

// 	// go func() {
// 	// 	for _, value := range <-greeting {
// 	// 		fmt.Println(value)
// 	// 	}
// 	// }()

// 	// for _, value := range <-greeting {
// 	// 	fmt.Println(value)
// 	// }

// 	// receive := <-greeting

// 	// fmt.Println(receive)

// 	buffedChan()

//		fmt.Println("End of program")
//	}
func BuffedReceiveOnlyChan() {
	ch := make(chan int, 2)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 2
	}()
	fmt.Println("Value:", <-ch)
	fmt.Println("Value:", <-ch)
}

// buffered channels
func buffedChan() {
	fmt.Println("==== Blocking on SEND ONLY IF THE BUFFER IS FULL =====")

	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println("Receiving from buffer")

	go func() {
		fmt.Println("Goroutine 2 second timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
		fmt.Println("Received:", <-ch)
	}()
	ch <- 3
	fmt.Println("Buffered Channels")
}
