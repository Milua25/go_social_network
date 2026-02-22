package LinkedList

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

// NewNode Instantiate a new Node
func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

// SingleLinkedList Single Linked List
type SingleLinkedList struct {
	Head *Node
}

func (sl *SingleLinkedList) Append(data int) {
	node := NewNode(data)

	if sl.Head == nil {
		sl.Head = node
	} else {
		current := sl.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

// RemoveData
func (sl *SingleLinkedList) RemoveData(data int) {
	if sl.Head == nil {
		return
	}

	if sl.Head.Data == data {
		sl.Head = sl.Head.Next
		return
	}

	current := sl.Head

	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Transverse to display each node
func (sl *SingleLinkedList) Transverse() {
	current := sl.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
	fmt.Println("End of List")
}

type DoubleNode struct {
	Data int
	Prev *DoubleNode
	Next *DoubleNode
}

func NewDoubleNode(data int) *DoubleNode {
	return &DoubleNode{Data: data, Prev: nil, Next: nil}
}

type DoubleLinkedList struct {
	Head *DoubleNode
}

func (dl *DoubleLinkedList) InsertAtEnd(data int) {
	node := NewDoubleNode(data)

	if dl.Head == nil {
		dl.Head = node
	} else {
		current := dl.Head

		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
		node.Prev = current
	}
}

func (dl *DoubleLinkedList) InsertAtStart(data int) {

	if dl.Head == nil {
		return
	}

	if dl.Head.Next == nil {
		dl.Head = nil
		return
	}

	newHead := dl.Head.Next

	newHead.Prev = nil

	dl.Head = newHead.Next

}

func (list LinkedList) Reverse() {
	var prev *Node = nil

	var current *Node = list.Head

	var next *Node = nil

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	list.Head = prev
	
}
