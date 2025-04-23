package main

import "fmt"

var (
	ErrUserRepoyNotFound = fmt.Errorf("user repository not found")
)

type ErrDivideByZero struct {
	IntA int
	IntB int
	Msg  string
}

func (e ErrDivideByZero) Error() string {
	return fmt.Sprintf("divide by zero: %s (a = %d, b = %d)", e.Msg, e.IntA, e.IntB)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide: %w", ErrDivideByZero{
			IntA: a,
			IntB: b,
			Msg:  "cannot divide by zero",
		})
	}
	return a / b, nil
}

func UserRepoAction(p string) (string, error) {
	if p == "" {
		return p, fmt.Errorf("user repo action: %w", ErrUserRepoyNotFound)
	}
	return p, nil
}

func UserUCAction(p string) (string, error) {
	q, err := UserRepoAction(p)
	if err != nil {
		return q, fmt.Errorf("user use case action: %w", err)
	}
	return q, nil
}

func UserHandlerAction(p string, a int, b int) (string, int, error) {
	q, err := UserRepoAction(p)
	if err != nil {
		return q, 0, fmt.Errorf("user handler action: %w", err)
	}

	d, err := divide(a, b)
	if err != nil {
		return q, 0, fmt.Errorf("user handler action: %w", err)
	}

	return q, d, nil
}

func main() {
	fmt.Println("Error Handling in Go")

	_, _, err := UserHandlerAction("", 1, 2)
	if err != nil {
		fmt.Println("Error 1:", err)
	}

	_, _, err = UserHandlerAction("tyty", 1, 0)
	if err != nil {
		fmt.Println("Error 2:", err)
	}
}
