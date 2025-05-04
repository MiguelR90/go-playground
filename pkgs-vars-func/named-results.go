package main

import "fmt"

// return values can be names useful for documenting the meaning of the return value
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// returns without variable default to variables defined in the func signature
	// this are known as "naked" returns should only be used in short functions
	// this can affect the readability of programs
	return

}

func main() {
	fmt.Println(split(100))
}
