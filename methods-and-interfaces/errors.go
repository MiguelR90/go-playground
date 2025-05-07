package main

import (
	"fmt"
	"time"
)

// go expresses error states as values
// the error type is builtin interface similar to fmt.Stringer
//
//	type error intereface {
//	    Error() string
//	}
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	// functions often return error value and calling code should handle errors
	// by testing whether error equals nil

	// nil error denotes no error
	// a non-nil error denotes failure
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
