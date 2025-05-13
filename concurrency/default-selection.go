package main

import (
	"fmt"
	"time"
)

// NOTE: select is not just use in the channel context?
// make you to follow up on select usecases / what select is trying to solve for

func main() {
	// hmm maybe not tick and after return channels???
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
