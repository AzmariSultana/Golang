package main
import "fmt"

func main() {

	var t any // "any" can hold a value of any type; unassigned, so it is nil
	fmt.Println("Hello, World!") // Println adds a newline automatically

	fmt.Printf("Hello " + "World\n") // "+" joins strings; Printf doesn't add a newline, so "\n" is needed

	fmt.Printf("Hello %s!\n", "World") // %s = placeholder for a string

	fmt.Printf("%d\n", 1+1) // %d = placeholder for an integer

	fmt.Println(1+1)

	fmt.Println(true, false)

	fmt.Printf("%+v\n", []int{1, 2, 3}) // %v prints any value in default format (%+v also shows struct field names); []int{...} is a slice

	fmt.Printf("%+v\n", t) // prints <nil>

}
