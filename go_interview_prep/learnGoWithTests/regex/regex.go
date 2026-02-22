package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said: \"I am great\"")
	fmt.Println(`He said: "He is a boy"`)

	// Compile a regular expression.
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "invalidEmail"

	fmt.Println(re.MatchString(email1))
	fmt.Println(re.MatchString(email2))

	// Capture groups
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	//date := "2006-01-02"

	fmt.Println(re.FindStringSubmatch("2006-01-02"))
	fmt.Println(re.FindStringSubmatch("2006-13-02")[0])
	fmt.Println(re.FindStringSubmatch("20-0001-32"))
	fmt.Println(re.FindStringSubmatch("2006-01-02T15:04:05Z"))
}
