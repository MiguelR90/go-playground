package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// this go routine will do both send and receive messages
	go func() {
		// receives 10 messages from channel c and print them to stdout
		for range 10 {
			fmt.Println(<-c)
		}
		// once loop is done we push a 0 message into the quit channel
		quit <- 0
	}()

	// this is the
	fibonacci(c, quit)
}
