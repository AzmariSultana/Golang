package main

import "fmt"

// Go has no enum keyword; enums are made with a const block and iota.
// iota starts at 0 and increases by 1 on each line of the block.
const (
	Friday = iota + 1 // starts from 1 instead of 0; following lines repeat "iota + 1"
	Saturday
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
)

type LogLevel int // new named type based on int, for type safety

const (
	LogError LogLevel = iota // iota resets to 0 in a new const block; all below get type LogLevel
	LogWarn
	LogInfo
	LogDebug
	LogFatal
)

func main() {
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)

}
