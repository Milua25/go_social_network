package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func main() {

	//x, y := 1, 2
	//x, y = swap(x, y)
	//fmt.Println(x, y)
	//
	//x1, y1 := "John", "Smith"
	//x1, y1 = swap(x1, y1)
	//fmt.Println(x1, y1)
	elements := Stack[int]{}
	elements.push(4)
	elements.push(3)
	elements.push(2)
	elements.push(1)
	//fmt.Println(elements.items)
	elements.printAll()
	elements.pop()
	elements.pop()
	elements.pop()
	elements.pop()
	fmt.Println(elements.isEmpty())

	// rem
}

func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) pop() (T, bool) {
	elementToBeDeleted := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return elementToBeDeleted, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.items) == 0
}

func swap[T any](a, b T) (T, T) {
	return b, a
}

func (s *Stack[T]) printAll() {
	if len(s.items) == 0 {
		fmt.Println("empty stack")
		return
	}

	fmt.Println(s.items)
	for _, item := range s.items {
		fmt.Println(item)
	}
}
