package main
import "fmt"

func main() {
	names := []string{"Azmari", "Sultana", "Tushmi"} // slice literal: like an array, but no length inside []
	fmt.Printf("Names: %v\n", names)

	items := make([]int, 3, 5) // make(type, length, capacity): 3 zeros, room for 5

	fmt.Printf("Items: %v, Length: %d, Capacity: %d\n", items, len(items), cap(items))
	// append adds to the end and returns the updated slice, so assign it back
	items = append(items, 1)
	items = append(items, 2) // length 5: capacity 5 is now full
	items = append(items, 3) // no room left: Go allocates a bigger array (capacity 10) and copies the elements
	items = append(items, 4)
	fmt.Printf("Items: %v, Length: %d, Capacity: %d\n", items, len(items), cap(items))

	fmt.Printf("%v\n", items[3:7]) // elements from index 3 up to, but not including, index 7
}
