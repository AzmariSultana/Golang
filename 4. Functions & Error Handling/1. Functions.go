package main
import "fmt"

// func name(parameter type) { body }; no return type, so it returns nothing
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// the type after the brackets is the return type
func add(a int, b int) int {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	return a + b // return sends the value back to the caller
}

func calculateArea(length float64, width float64) float64 {
	if width <= 0 || length <= 0 { // check the input before using it
		fmt.Println("Error: Length and width must be positive numbers.")
		return 0
	}
	return length * width
}

func main() {
	greet("Azmari Wonderland")
	add(1, 2) // the returned value is ignored here, which is allowed

	area := calculateArea(5, 10)
	fmt.Printf("Area: %f\n", area) // %f prints 6 decimal places by default
}
