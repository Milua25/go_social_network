package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Please enter your name and age: ")
	// _, err := fmt.Scan(&name, &age) // take input from a user
	//	_, err := fmt.Scanln(&name, &age)
	_, err := fmt.Scanf("%s, %d", &name, &age)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Printf("%s is %d years old\n", name, age)
}
