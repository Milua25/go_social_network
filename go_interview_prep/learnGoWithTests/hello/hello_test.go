package hello

import "testing"

/*
Rules for testing
It needs to be in a file with a name like xxx_test.go

The test function must start with the word Test

The test function takes one argument only t *testing.T

To use the *testing.T type, you need to import "testing", like we did with fmt in the other file
*/

func TestHello(t *testing.T) {

	// test 1
	//t.Run("saying hello to people", func(t *testing.T) {
	//	got := Hello("Chris")
	//	want := "Hello, Chris!"
	//	if got != want {
	//		t.Errorf("got %q, want %q", got, want)
	//	}
	//})

	// refactored test 1
	//t.Run("saying hello to people", func(t *testing.T) {
	//	got := Hello("Chris")
	//	want := "Hello, Chris!"
	//	assertCorrectMessage(t, got, want)
	//})

	// test 2
	//t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
	//	got := Hello("")
	//	want := "Hello, World!"
	//	assertCorrectMessage(t, got, want)
	//})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English", func(t *testing.T) {
		got := Hello("Elodie", "")
		want := "Hello, Elodie!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English", func(t *testing.T) {
		got := HelloSwitch("Elodie", "French")
		want := "Bonjour, Elodie!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
