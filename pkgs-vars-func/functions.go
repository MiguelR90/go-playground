package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

// when consecutive arguments share type it's okay to obmit from except last one
// this is equivalent as the above function that types each argument or function
// paramter
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(40, 2))
}
