package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// walk walks the tree t sending all values from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	// In a binary tree: Left.Value <= Value <= Right.Value
	// Thus, depth-first traversal to make sure we get values in order

	// how do you do this without recursion?
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	// start by sending current value
	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	go Walk(t1, ch1)

	ch2 := make(chan int, 10)
	go Walk(t2, ch2)

	var u, v int
	for range cap(ch1) {
		if u, v = <-ch1, <-ch2; u != v {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println("Check true ==", Same(tree.New(1), tree.New(1)))
	fmt.Println("Check false ==", Same(tree.New(1), tree.New(2)))
}
