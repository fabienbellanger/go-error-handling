package main

import (
	"fmt"

	go_error_handling "github.com/fabienbellanger/go-error-handling"
)

func main() {
	fmt.Println("Error Handling in Go")

	_, _, err := go_error_handling.UserHandlerAction("", 1, 2)
	if err != nil {
		fmt.Println("Error 1:", err)
	}

	_, _, err = go_error_handling.UserHandlerAction("tyty", 1, 0)
	if err != nil {
		fmt.Println("Error 2:", err)
	}
}
