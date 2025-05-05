package main

import (
	"errors"
	"fmt"

	"github.com/fabienbellanger/go-error-handling/xerr"
)

var ErrClassic = errors.New("a classic error")

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("\nError Handling in Go")
	fmt.Print("--------------------\n\n")

	testErr()
}

func testErr() {
	err := testB(false)
	fmt.Printf("Err: \n%s\n\n", err)

	e, _ := err.JSON()
	fmt.Printf("Err JSON: \n%s\n\n", e)
}

func testA(ok bool) (err xerr.Err) {
	if !ok {
		return xerr.NewErr(ErrClassic, "Error in testA", Person{Name: "test", Age: 12}, nil)
	}

	return
}

func testB(ok bool) xerr.Err {
	err := testA(ok)
	if err.IsError() {
		if err.Is(ErrClassic) {
			fmt.Println("Error in testB: a classic error")
		}

		return xerr.NewErr(errors.New("my super error"), "Error in testB", nil, &err)
	}

	return xerr.EmptyErr()
}
