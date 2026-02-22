package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum array of numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		// to unfurl an array into a slice [:]...
		got := Sum(numbers[:]...)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		// to unfurl a slice []...
		sum := Sum(numbers...)
		want := 15
		if sum != want {
			t.Errorf("got %d want %d given, %v", sum, want, numbers)
		}
	})

	t.Run("sum slice count", func(t *testing.T) {
		sum := SumAll([]int{1, 2, 3, 4, 5}, []int{6, 7, 8})
		want := []int{15, 21}
		if !slices.Equal(sum, want) {
			t.Errorf("got %d want %d\n", sum, want)
		}
	})

	t.Run("sum tails", func(t *testing.T) {
		sum := SumTails([]int{1, 2, 3, 4, 5}, []int{5, 6, 7})
		want := []int{14, 13}
		if !reflect.DeepEqual(sum, want) {
			t.Errorf("got %d want %d\n", sum, want)
		}
	})

	t.Run("sum empty slice", func(t *testing.T) {
		sum := SumTails([]int{}, []int{5, 6, 7})
		var want []int
		checkSums(t, sum, want)
	})

	t.Run("append to a slice", func(t *testing.T) {
		sl1 := []int{1, 2, 3}
		sl2 := []int{4, 5, 6}

		// update slice
		got := UpdateSlice(sl1, sl2)
		want := []int{1, 2, 3, 4, 5, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
