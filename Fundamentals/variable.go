package main
import "fmt"

func main() {
	var greeting string //zero_value is an empty string ""
	greeting = "Hello, World!"
	fmt.Println(greeting)

	var count int // zero value is 0
	count = 10
	fmt.Println(count)

	var isRunning bool // zero value is false
	isRunning = true
	fmt.Println(isRunning)

	var firstname, lastname string // multiple variables of the same type in one line
	firstname = "Azmari"
	lastname = "Sultana"
	fmt.Println(firstname, lastname)

	email := "azmari.sultana@example.com" //short declaration, type is inferred
	// ":=" declares and assigns together; only allowed inside functions
	fmt.Println(email)

	age:= 16 //short declaration, type is inferred
	fmt.Println(age)

	var year = 2025 // type (int) inferred from the value
	fmt.Println(year)
}
