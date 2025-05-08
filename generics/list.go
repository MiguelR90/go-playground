package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
GO also supports parameterized type ( not just functions )

A type can be parameterized with a type params which could be useful for
implementing generic data structures

Ex: a type declaration for a singly-linked list holding any type
*/

// NOTE: first time we're introduced to the dynammic type construct similar to py Any
type List[T any] struct {
	next *List[T]
	val  T
}

func (root *List[T]) All() func(yield func(i int, item T) bool) {
	return func(yield func(i int, item T) bool) {
		for i, current := 0, root.next; current != nil; i, current = i+1, current.next {
			if !yield(i, current.val) {
				return
			}
		}
	}
}

func (root *List[T]) Rappend(elem T) {
	// start at the root and move pointer until last element is reached
	// insert the new elem at the end
	current := root
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: elem}

}

func (root *List[T]) Lappend(elem T) {
	// adds element to beginning of the list
	// and updates the pointer reference
	newNode := &List[T]{val: elem, next: root.next}
	root.next = newNode
}

func (root *List[T]) String() string {

	// FIXME: can a builder be initalized?
	var builder strings.Builder
	builder.WriteString("[ ")

	// NOTE: Looping over a range iterator!! (new syntax since 1.23 in 2024)
	for _, val := range root.All() {
		builder.WriteString(fmt.Sprintf("elem=%v ", val))
	}

	builder.WriteString("]")
	return builder.String()
}

func (root *List[T]) Len() int {
	counter := 0
	for current := root.next; current != nil; current = current.next {
		counter++
	}
	return counter
}

// this function needs to raise and error
func (root *List[T]) Insert(i int, elem T) (*List[T], error) {
	if i < 0 {
		return nil, errors.New("index out of bounds")
	}

	// Insert at head
	if i == 0 {
		root.Lappend(elem)
		return root, nil
	}

	current := root.next
	for pos := 0; current != nil && pos+1 < i; pos++ {
		current = current.next
	}

	if current == nil {
		return nil, errors.New("index out of bounds")
	}

	newNode := &List[T]{val: elem, next: current.next}
	current.next = newNode

	return root, nil
}

func main() {

	// Add() adds elements to the beginning of the list
	// this returns double zero becuse we need the concept of a root node
	list := new(List[int])
	for i := 0; i < 10; i++ {
		list.Lappend(i)
	}
	fmt.Println(list.String())

	// Append() adds elements to the end of the list
	other := new(List[string])
	for i := 0; i < 10; i++ {
		other.Rappend(fmt.Sprintf("item%v", i))
	}
	fmt.Println(other.String())
	fmt.Println(other.Len())

	other, _ = other.Insert(10, "NEW")
	fmt.Println(other.String())

	other, _ = other.Insert(0, "NEW")
	fmt.Println(other.String())

	other, _ = other.Insert(5, "NEW")
	fmt.Println(other.String())

	_, err := other.Insert(18, "NEW")
	fmt.Println(err)

	_, err = other.Insert(-1, "NEW")
	fmt.Println(err)
}
