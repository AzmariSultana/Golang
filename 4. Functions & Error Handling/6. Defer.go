package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: deferred") // runs when simpleDefer returns, after all the Middle lines
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")

}

func lifoSimpleDefer() {
	fmt.Println("Function lifoSimpleDefer: Start")
	// deferred calls run last in, first out (LIFO): "Second" runs before "First"
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second: deferred")
	fmt.Println("Function lifoSimpleDefer: Middle")
}
func main() {
	file, err := os.Create("./defer.txt") // creates (or empties) defer.txt in the folder you run from
	if err != nil {
		fmt.Println(err)
		return // file is nil, so there is nothing to close
	}
	defer file.Close() // common pattern: close the file when main ends, right after opening it
	//simpleDefer()
	lifoSimpleDefer()

	fmt.Println("Last in main()")

}
