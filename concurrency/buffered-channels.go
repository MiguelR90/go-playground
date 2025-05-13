package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// since channel is buffered this will result in a dead lock
	// also known as overfilling the buffer
	// unclear why something like this would be useful
	// ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// trying to pull from an empty? channel also results in a deadlock error
	// fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-ch)
}
