package main

import "fmt"

// channels are a type condiut (tube or channel for direct communication)
// through which you can send and receive values with the channel operator <-
// <- channel operator: data flows in the direction of the arrow

// like maps and slices, channels must be created before use
// ch := make(chan, int)
// by default, sends and receives block until the other side is ready
// this allows goroutines to sync without explicit locks or condition vars

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	// split the slice in two
	// distributed the sum calculation between two go routines
	// retrieve result and combine
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	// this is more or less a map reduce operation
	fmt.Println(x, y, x+y)
}
