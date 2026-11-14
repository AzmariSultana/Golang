package main

import "fmt"

type LogLevel int // custom type for log levels, based on int

// iota numbers the levels from 0: Trace = 0, Debug = 1, ... Error = 4
const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

// names in the same order as the constants, so levelNames[level] gives the right name
var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

// String is a method on LogLevel that returns the level's name.
// Because of this method, fmt can print a LogLevel as text (for example with %s or %v).
func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError { // out of range: avoid an index out of range panic
		return "Unknown"
	}

	return levelNames[l] // the level's number is used as the slice index
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log level: %d %s\n", level, level.String()) // %d prints the number, %s the name
}
func main() {

	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(10) // allowed: the untyped constant 10 becomes a LogLevel; String() returns "Unknown"

}
