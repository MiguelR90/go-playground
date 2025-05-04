package main

import "fmt"

func main() {

	// there is no while keyword in go so we can get a while construct by
	// omitting the init and post statement from the for loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
