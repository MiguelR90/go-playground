package main

import "fmt"

func main() {
	sum := 0

	// WARNING: lsp alerts that there is newer syntax for the for loop
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i, sum)
	}
	fmt.Println(sum)

}
