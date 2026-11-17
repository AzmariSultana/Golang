package main

import "fmt"

// recursion: factorial calls itself with a smaller number each time
func factorial(n int) int {
	if n <= 1 { // base case: stops the recursion
		return 1
	}
	return n * factorial(n-1)
}

// returns a function; the returned function "closes over" i and keeps it between calls
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	fmt.Println(factorial(5)) // 5 * 4 * 3 * 2 * 1 = 120

	nextInt := intSeq() // nextInt has its own i, starting at 0
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// anonymous function (a function without a name) stored in a variable
	logger := func(msg string) {
		fmt.Println(msg)
	}

	logger("Hello World")
	logger("Hello World")

}
