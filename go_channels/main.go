package main

import "fmt"

func main() {
	// unbuffered channel
	uc := make(chan int)

	// buffered channel
	bc := make(chan int, 1)

	c := make(chan int, 1)

	// send only channel
	sc := make(chan<- int, 1)

	//c <- 42 // sending to a channel this will cause a block

	go func() {
		uc <- 42 // sending to a channel
	}()

	go func() {
		bc <- 45 // sending to a channel
	}()

	go func(num int) {
		sc <- num
	}(2)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("UnBufferred Channel:", <-uc)
	fmt.Println("Bufferred Channel:", <-bc)
	// fmt.Println("Recieve Only Channel:", <-rc)
}

// func bar(rc <-chan int) {
// 	fmt.Println(<-rc)
// }

func foo(sc chan<- int) {
	for i := 0; i < 101; i++ {
		sc <- i
	}
	close(sc)
}
