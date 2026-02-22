package main

import (
	"fmt"
	"time"
)

func main() {
	// Non Blocking Receive

	ch := make(chan int)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No messages available")
	}

	// Non Blocking Send operation
	select {
	case ch <- 1:
		fmt.Println("Sent message")
	default:
		fmt.Println("Channel is full")
	}

	// Non Blocking Operations in realtime
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Received:", d)
			case <-quit:
				fmt.Println("Quitting...")
				return
			default:
				fmt.Println("No data received")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
}
