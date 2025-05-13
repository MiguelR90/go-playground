package main

import "fmt"

// NOTE: only the sender should close a channel, never the receiver.
// sending on a closed channel with cause a panic

// NOTE: channels aren't like files that need to be closed afterwards
// closing is only necessary when the receiver must be told there are
// no more values coming and therefore terminate a range loop
// this causes a coupling between the sender and ther receiver!!!

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	// this forloop syntax can be improved with new range int
	// for i := 0; i < n; i++ {
	for range n {
		c <- x
		x, y = y, x+y

	}
	// sender can close a channel to indicate that no more values will be sent
	// receivers can test whether a channel has been closed by assigning a second
	// param to the receiver expression: after
	// v, ok := <-ch
	close(c)
}

func main() {

	// FIXME: i am still not quite comfortable with the make syntax; when it's appropriate and they it is not
	c := make(chan int, 25)
	go fibonacci(cap(c), c)

	// the loop for i := range c receives values from the channel repeatedly until it is closed.
	// range func likely uses v, ok := <-ch syntax under the hood
	for i := range c {
		fmt.Println(i)
	}
}
