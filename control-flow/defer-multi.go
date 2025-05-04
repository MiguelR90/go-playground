package main

import "fmt"

func main() {
	fmt.Println("start: counting")
	// deferred function calls are pushed onto a stack and are executed
	// in LIFO (last in first out) order.
	// this will print the numbers backwards from 9, 8, 7 ... 1, 0
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done: counting")
}
