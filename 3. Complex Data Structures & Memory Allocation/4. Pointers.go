package main
import "fmt"

// val is a copy of the caller's value, so changing it does not affect the caller
func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}


// *int is a pointer to an int: it holds the memory address of the caller's variable
func modifyPointer(val *int){
	if val == nil { // a nil pointer points to nothing; dereferencing it would crash
		fmt.Println("value is nil")
		return
	}
	*val = *val * 10 // dereferencing the pointer to modify the value it points to
	fmt.Printf("modifyPointer: %+v\n", *val)
}


func main() {
	num := 10
	modifyValue(num) // passes a copy of num
	fmt.Printf("num: %+v\n", num) // still 10

	modifyPointer(&num) // & gives the address of num
	fmt.Printf("num: %+v\n", num) // now 100

	grade := 50
	gradePointer := &grade
	fmt.Printf("gradePtr: %+v\n", gradePointer) // prints the address, such as 0xc00000a0f8
	fmt.Printf("gradePointer: %+v\n", *gradePointer) // * reads the value at that address: 50
}
