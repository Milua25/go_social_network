package main

import "fmt"

func adder() func() int {
	i := 0
	return func() int {
		i++
		fmt.Printf("added + 1 to %d\n", i)
		return i
	}
}

func main() {

	//add := adder()
	//println(add())
	//println(add())
	//println(add())
	//println(add())
	//
	//add2 := adder()
	//println(add2())
	//println(add2())

	newClosureFunc := func() func(x int) int {
		countDown := 99
		return func(b int) int {
			countDown -= b
			return countDown
		}
	}
	azz := newClosureFunc()
	println(azz(2))
	println(azz(2))
	println(azz(2))
	println(azz(2))
	println(azz(2))
	println(azz(2))
}
