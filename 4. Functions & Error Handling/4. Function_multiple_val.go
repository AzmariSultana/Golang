package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// returns two values: the result and an error (nil means no error)
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero") // errors.New creates a simple error with a message
	}
	return a / b, nil
}

// named return values: firstName and lastName are declared here and start as ""
func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ") // "Azmari Sultana" -> ["Azmari", "Sultana"]
	firstName = parts[0]
	lastName = parts[1] // panics if the name has no space, because parts has only 1 element

	return // "naked" return: returns firstName and lastName
}

func main() {

	time.Now() // does nothing here: the current time is returned but not used

	value, err := divide(10, 0)
	if err != nil { // always check the error before using the value
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	firstName, lastName := splitName("Azmari Sultana")

	fmt.Println(firstName, lastName)
}
