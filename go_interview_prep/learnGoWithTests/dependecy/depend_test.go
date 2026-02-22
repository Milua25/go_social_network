package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {

	t.Run("prints a greeting using bytes.Buffer", func(t *testing.T) {
		buffer := new(bytes.Buffer)

		Greeting(buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris!\n"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	//t.Run("prints a greeting using fmt.Println", func(t *testing.T) {
	//
	//})

}
