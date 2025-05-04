package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("GO runs on ")

	// switch is shorthand for a sequence of if else statements
	// it runs the first case whose value is equal to the conditional express

	// runs the selected case only, not all cases that follow
	// until java, c, php etc. the break statement is not needed
	// cases don't have to be constant and values don't have to be integer
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("linux.")
	default:
		// some other operating system
		// freebsd, openbsd and windows
		fmt.Printf("%s.\n", os)
	}

}
