package main

import (
	"fmt"
	"strings"
)

// custom error type that records which operation failed and with which inputs
type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

const (
	division       = "Division"
	divisionErrMsg = "Division by zero is not allowed"
)

// Error() makes MathError satisfy the error interface
func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA)) // Sprintf returns the text instead of printing it
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s",
		e.Operation,
		strings.Join(inputs, ","), // ["a=10", "b=0"] -> "a=10,b=0"
		e.Message)
}

func Sum(numbers ...int) int {
	defer fmt.Println("Sum finished") // printed when Sum returns, before the caller prints the total

	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SafeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{ // a pointer works as an error too, because Error() has a value receiver
			Operation: division,
			InputA:    a,
			InputB:    b,
			Message:   divisionErrMsg,
		}
	}

	return a / b, nil
}

func main() {

	fmt.Println(Sum(1, 2, 3))

	value, err := SafeDivision(10, 0)
	if err != nil {
		fmt.Println(err) // Println calls err.Error() to get the message
	}

	fmt.Println("SafeDivision", value) // value is 0 because the division failed

}
