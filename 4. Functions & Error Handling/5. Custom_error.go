package main

import (
	"errors"
	"fmt"
	"time"
)

// sentinel errors: fixed error values that callers can check for; names start with Err by convention
var ErrDivisionByZero = errors.New("division by zero")
var ErrNumTooLarge = errors.New("number too large")

// a custom error type can carry extra details about what went wrong
type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

// any type with an Error() string method satisfies the built-in error interface
func (op OpError) Error() string {
	return op.Message
}

// constructor function: builds an OpError and returns a pointer to it
func NewOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

// return type is error, so the *OpError is returned as an error (not called in main)
func DoSomething() error {
	return NewOpError("doSomething", 100, "do something failed", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if a > 1000 {
		return 0, ErrNumTooLarge
	}

	return a / b, nil
}

func main() {

	value, err := divide(1001, 1) // 1001 > 1000, so ErrNumTooLarge is returned
	if err != nil {
		if errors.Is(err, ErrDivisionByZero) { // errors.Is checks whether err is (or wraps) that error
			fmt.Println("divide by zero")
		} else if errors.Is(err, ErrNumTooLarge) {
			fmt.Println("number too large")
		}
		return // stop main here; value is not valid
	}

	fmt.Println(value)
}
