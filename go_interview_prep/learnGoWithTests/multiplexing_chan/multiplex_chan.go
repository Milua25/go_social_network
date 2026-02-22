package main

import (
	"fmt"
	"time"
)

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- 1
// 		time.Sleep(2 * time.Second)
// 		ch2 <- 1
// 	}()

// 	// condition
// 	select {

// 	case msg1 := <-ch1:
// 		fmt.Println("Receive from ch1:", msg1)

// 	case msg2 := <-ch2:
// 		fmt.Println("Receive from ch2:", msg2)

// 	default:
// 		fmt.Println("No channel")
// 	}
// }

// implement timeout in channels
func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		defer close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel Close!!!")
				return
			}
			fmt.Println(msg)

			// case <-time.After(5 * time.Second):
			// 	fmt.Println("Timeout!!!")
		}
	}

}
