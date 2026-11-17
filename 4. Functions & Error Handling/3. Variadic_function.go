package main

import "fmt"

// ...int accepts any number of ints; inside the function, numbers is a []int
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func config(numbers ...int) {
	if len(numbers) > 0 { // called with no arguments, numbers is empty
		first := numbers[0]
		fmt.Println("First number:", first)
	} else {
		fmt.Println("Default number")
	}
}

func main() {

	fmt.Println(sum(1, 2, 3, 4))

	config() // no arguments, so the else branch runs

}
