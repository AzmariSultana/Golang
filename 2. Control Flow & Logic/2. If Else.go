package main
import "fmt"

func main() {
	temp := 25
	if temp > 30 { // no parentheses around the condition, but the braces are required
		fmt.Println("It's hot!")
	} else { // else must be on the same line as the closing brace "}"
		fmt.Println("It's not hot.")
	}

	score := 85
	// conditions are checked from top to bottom; only the first true block runs
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Failed")
	}

	// map of username to access permission; map[key type]value type
	userAccess := map[string]bool{
		"azmari": true,
		"sultana": false,

	}
	// short statement before the condition; hasAccess and ok only exist inside this if/else
	// ok is true if the key is in the map; && means "and"
	if hasAccess, ok := userAccess["sultana"]; ok && hasAccess {
		fmt.Println("Sultana can access the system.")
	} else {
		fmt.Println("Sultana cannot access the system.")
	}
}
