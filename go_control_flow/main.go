package main

import "fmt"

func init() {
	fmt.Println("I run before the main program 😀!!!")
}

func main() {
	fmt.Println("Control Flow!!!")
	x := 50
	switch {
	case x < 42:
		fmt.Println("X is less than 42")
	default:
		fmt.Println("This is the default response")
	}

	switch x {
	case 45:
		fmt.Println("x is 45")

	default:
		fmt.Println("x is not 45")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer loop %v \t Inner loop %v\n", i, j)
		}
	}

	// Maps

	m := map[string]string{
		"FirstName": "Ayo",
	}
	for i, v := range m {
		fmt.Println(i, v)
	}
}
