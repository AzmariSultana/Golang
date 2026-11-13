package main

import "fmt"
const HOST string = "localhost" // package-level constant; const values can never be changed

func main() {
	println(HOST) // built-in println, mainly for debugging; fmt.Println is preferred

	const AppName string = "Go" // local constant, only visible inside main()
	fmt.Println(AppName)

	const pi float64 = 3.1415926 // float64 = 64-bit decimal, more precise (default for decimals)
	fmt.Println(pi)

	const rate float32 = 5.2 // float32 = 32-bit decimal, less memory, less precise
	fmt.Println(rate)

}
