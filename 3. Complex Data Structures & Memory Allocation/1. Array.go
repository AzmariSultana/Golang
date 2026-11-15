package main

import "fmt"

func main() {
	var numbers [2]int // array of exactly 2 ints; the length is part of the type
	fmt.Printf("%v\n", numbers) // every element starts at the zero value: [0 0]

	numbers[0] = 1 // positions (indexes) start at 0
	numbers[1] = 2
	fmt.Printf("%+v\n", numbers)

	primes := [4]int{2, 3, 5, 7} // array literal: declare and fill in one step
	fmt.Printf("%+v\n", primes)
	primes[3] = 11 // replace the last element
	fmt.Printf("%+v\n", primes)

	for i := 0; i < len(primes); i++ { // len gives the number of elements
		fmt.Printf("%d\n", primes[i])
	}

	var matrix [2][3]int // 2 rows, each an array of 3 ints
	matrix[0][0] = 1 // [row][column]
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	fmt.Printf("%+v\n", matrix)
}
