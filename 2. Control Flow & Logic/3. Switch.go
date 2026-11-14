package main
import "fmt"
import "time"

func main() {

	day := "Friday"
	fmt.Println("Today is", day)

	switch day { // compares day with each case from top to bottom
	case "Friday", "Saturday": // one case can match several values, separated by commas
		fmt.Println("It's the weekend!") // no break needed; Go stops after the matching case
	case "Monday", "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Work days. Lots of work to do.")
	default: // runs when no case matches
		fmt.Println("Invalid day.")
	}

	// no value after "switch", so each case is a condition, like an if-else chain
	switch hour := time.Now().Hour(); { // hour (0 to 23) only exists inside this switch
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// a function stored in a variable; interface{} is the same as any, so i can hold any type
	checkType := func(i interface{}) {
		switch v := i.(type) { // type switch: matches the type of i; v holds the value as that type
		case int:
			fmt.Println("Integer:", v)
		case string:
			fmt.Println("String:", v)
		case bool:
			fmt.Println("Boolean:", v)
		default:
			fmt.Println("Unknown type")
		}
	}

	checkType(47)
	checkType("Hello")
	checkType(true)
	checkType(3.69) // 3.69 is a float64, which has no case, so default runs
}
