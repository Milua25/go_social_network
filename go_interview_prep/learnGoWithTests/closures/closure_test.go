package main

import (
	"reflect"
	"runtime"
	"testing"
)

func TestAdder(t *testing.T) {

	got := adder()

	want := func() int { return 1 }

	f1 := runtime.FuncForPC(reflect.ValueOf(got).Pointer()).Name()
	f2 := runtime.FuncForPC(reflect.ValueOf(want).Pointer()).Name()

	if f1 != f2 {
		t.Errorf("got %v want %v", f1, f2)
	}
}
