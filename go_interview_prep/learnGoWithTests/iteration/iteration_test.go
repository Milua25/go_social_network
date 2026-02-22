package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected %q, got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// setup
	for b.Loop() {
		// code to measure...
		Repeat("a", 10)
	}
	//..cleanup..

}

func ExampleRepeat() {
	result := Repeat("b", 5)
	fmt.Println(result)
	// Output: bbbbb
}
