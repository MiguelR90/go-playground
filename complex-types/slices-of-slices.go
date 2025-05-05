package main

import (
	"fmt"
	"strings"
)

func main() {
	// slices can contain any type including other slices
	// this represets a 2 x 2 array of string, ie., a tick tack toe board

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	// loop over the rows printing one at a time
	// using strings.Join to join together the inner slice of strings
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
