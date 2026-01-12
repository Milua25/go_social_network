package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/GoesToEleven/puppy"
)

type dog struct {
	first string
}

func main() {
	printRuntimeInfo()
	fmt.Println("Hello, World!")
	fmt.Println(puppy.Bark())
	puppy.From11()

	runWithWaitGroups()
	runDogDemo()

	fmt.Println(addT(1.0, 2.3))
	fmt.Println(addT(1, 2))
}

func printRuntimeInfo() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I am running!!")
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I am walking!!")
}

type young interface {
	walk()
	run()
}

func youngRun(y young) {
	y.run()
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | float64](a, b T) T {
	return a + b
}

func runDogDemo() {
	d1 := dog{"Henry"}
	d2 := &dog{"Puppy"}

	d1.walk()
	d1.run()
	d2.walk()
	d2.run()
	youngRun(d2)
}

func runWithWaitGroups() {
	const goroutines = 100

	var counter int64
	var workers sync.WaitGroup
	var announcer sync.WaitGroup

	workers.Add(goroutines)
	announcer.Add(1)

	go func() {
		defer announcer.Done()
		fmt.Println("I am a boy!!!")
	}()

	for i := 0; i < goroutines; i++ {
		go func() {
			defer workers.Done()
			//runtime.Gosched()
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter: ", atomic.LoadInt64(&counter))
		}()
	}

	workers.Wait()
	fmt.Println("Counter:", counter)
	fmt.Println("GoRoutines:", runtime.NumGoroutine())
	fmt.Println(counter)
	announcer.Wait()
}
