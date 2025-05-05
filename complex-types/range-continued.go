package main

import "fmt"

func main() {
	pow := make([]int, 10)

	// you can skip the index or the value by assigning an _
	// if you want only the index you can omit the second variable
	for i := range pow {
		// byte-wise shift (like 1e1,1e2, 1e3, ... except in base 2)
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
