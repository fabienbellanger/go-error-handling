// Package to test error handling in Go
//
// This package demonstrates how to handle errors in Go using the
// standard library's error handling capabilities.
package go_error_handling

import "fmt"

var (
	// ErrUserRepoyNotFound is an error that indicates a user repository was not found
	ErrUserRepoyNotFound = fmt.Errorf("user repository not found")
)

// ErrDivideByZero is a custom error type that indicates a division by zero error
type ErrDivideByZero struct {
	IntA int
	IntB int
	Msg  string
}

func (e ErrDivideByZero) Error() string {
	return fmt.Sprintf("divide by zero: %s (a = %d, b = %d)", e.Msg, e.IntA, e.IntB)
}

// Divide function that returns an error if the second argument is zero
//
// # Example:
//
//	result, err := Divide(10, 0)
//
// - This function demonstrates how to handle errors in Go.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide: %w", ErrDivideByZero{
			IntA: a,
			IntB: b,
			Msg:  "cannot divide by zero",
		})
	}
	return a / b, nil
}

// UserRepoAction simulates a user repository action
func UserRepoAction(p string) (string, error) {
	if p == "" {
		return p, fmt.Errorf("user repo action: %w", ErrUserRepoyNotFound)
	}
	return p, nil
}

// UserUCAction simulates a user use case action
func UserUCAction(p string) (string, error) {
	q, err := UserRepoAction(p)
	if err != nil {
		return q, fmt.Errorf("user use case action: %w", err)
	}
	return q, nil
}

// UserHandlerAction simulates a user handler action
func UserHandlerAction(p string, a int, b int) (string, int, error) {
	q, err := UserRepoAction(p)
	if err != nil {
		return q, 0, fmt.Errorf("user handler action: %w", err)
	}

	d, err := Divide(a, b)
	if err != nil {
		return q, 0, fmt.Errorf("user handler action: %w", err)
	}

	return q, d, nil
}
