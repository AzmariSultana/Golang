package main
import "fmt"
func main() {

	// map[key type]value type; each name (key) maps to a grade (value)
	studentGrades := map[string]int{
		"Azmari": 90,
		"Sultana": 85,
		"Tushmi": 60,
	}
	fmt.Printf("Student Grades: %v\n", studentGrades) // fmt prints maps sorted by key
	studentGrades["Azmari"] = 95 // the key already exists, so its value is replaced
	fmt.Printf("Student Grades: %v\n", studentGrades)

	azmari, ok := studentGrades["Azmari"] // ok is true if the key exists
	if ok {
		fmt.Printf("Azmari's grade: %d\n", azmari)
	} else {
		fmt.Println("Azmari's grade not found")
	}

	key := "Hossain"
	// !ok means the key is missing; value is then the zero value of int, 0
	if value, ok := studentGrades[key]; !ok {
		fmt.Printf("%s: %+v\n", key, value)
	}
	delete(studentGrades, "Tushmi") // removes the key; does nothing if it is not there
	fmt.Printf("Student Grades after deletion: %v\n", studentGrades)

	configs := make(map[string]int) // an empty map that is ready to use
	fmt.Printf("Configs: %v %T\n", configs, configs) // %T prints the type

	if configs == nil { // false: make creates a real map; only "var configs map[string]int" is nil
		fmt.Println("Configs is nil")
	}
}
