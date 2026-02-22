package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Printf("computing factorial of %d\n", n)
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	fmt.Printf("computing Fibonacci of %d\n", n)
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(factorial(5))

	fmt.Println(fibonacci(10))

	// practical use cases
}
