package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// done := make(chan struct{})

	// go func() {
	// 	fmt.Println("working...")
	// 	time.Sleep(2 * time.Second)
	// 	done <- struct{}{}
	// }()

	// <-done
	// fmt.Println("Finished")

	// chan2()

	//synchronizeData()

	ch := sendOnlyChan()

	receiveData(ch)
}

// function
func chan1() {
	ch := make(chan int)

	go func() {
		ch <- 9
		fmt.Println("sent value")
	}()

	value := <-ch
	fmt.Println(value)
}

func chan2() {
	numGoRoutines := 3

	done := make(chan int, 3)

	for i := range numGoRoutines {
		go func(id int) {
			fmt.Printf("GoRoutine %d working...\n", id)
			time.Sleep(time.Second)
			done <- id
		}(i)
	}

	for range numGoRoutines {
		// wait for each goRoutine to finish
		val := <-done

		fmt.Println(val)
	}

	fmt.Println("All GoRoutines are finished!!")

}

// ==== SYNCHRONIZING DATA EXCHANGE ====
func synchronizeData() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			str := "hello" + " " + strconv.Itoa(i)
			data <- str
			time.Sleep(100 * time.Millisecond)
		}
		defer close(data)
	}()

	for val := range data {
		fmt.Println("Received value:", val)
	}
}

// SEND ONLY CHANNEL

func sendOnlyChan() chan int {
	ch := make(chan int)

	go func(ch chan<- int) {
		for i := range 5 {
			ch <- i
		}
		defer close(ch)
	}(ch)

	return ch
}

// Receive Only Channel
func receiveData(ch <-chan int) {
	for i := range ch {
		fmt.Println("Received: ", i)
	}
}
