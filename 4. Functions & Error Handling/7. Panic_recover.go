package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong in mightPanic") // stops normal execution and starts unwinding the stack
	}

	fmt.Println("This function executed without a panic") // never reached when shouldPanic is true
}

func recoverable() {
	// recover only works inside a deferred function
	defer func() {
		if r := recover(); r != nil { // r is the value passed to panic
			fmt.Println("Recovered from panic:", r)
		}
	}()

	mightPanic(true)
}

func main() {

	recoverable() // the panic is recovered inside, so the program ends normally

}
