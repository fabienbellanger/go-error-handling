package main

import (
	"errors"
	"fmt"
	"log"

	go_error_handling "github.com/fabienbellanger/go-error-handling"
	"github.com/fabienbellanger/go-error-handling/apperror"
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

	err2 := testB(false)
	e, _ := err2.JSON()
	log.Printf("Err: %s", e)
}

func testA(ok bool) apperror.Err {
	if !ok {
		return apperror.NewErr(errors.New("a classic error"), "Error in testA", nil, nil)
	}
	return apperror.EmptyErr()
}

func testB(ok bool) apperror.Err {
	err := testA(ok)
	if err.IsErr() {
		return apperror.NewErr(nil, "Error in testB", nil, &err)
	}
	return apperror.EmptyErr()
}
