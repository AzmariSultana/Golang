package main
import "fmt"
import "slices" // standard package with helper functions for slices

func main() {

	fmt.Printf("---- Advanced Slicing Operations ----\n")
	original := []int{0,1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("original: %+v, length: %d, capacity: %d\n", original, len(original), cap(original))

	// slicing does not copy: s1 to s4 all share original's underlying array
	s1 := original[2:5] // low:high (excluding high)
	fmt.Printf("s1: %v, length: %d, capacity: %d\n", s1, len(s1), cap(s1)) // capacity counts from index 2 to the end: 8

	s2 := original[:4] // low:high (excluding high); low left out means 0
	fmt.Printf("s2 (original[:4]): %v, length: %d, capacity: %d\n", s2, len(s2), cap(s2))

	s3 := original[6:] // high left out means up to the end
	fmt.Printf("s3 (original[6:]): %v, length: %d, capacity: %d\n", s3, len(s3), cap(s3))

	s4 := original[:] // the whole slice
	fmt.Printf("s4 (original[:]): %v, length: %d, capacity: %d\n", s4, len(s4), cap(s4))

	fmt.Printf("s4 contains 8: %v\n", slices.Contains(s4, 8)) // true if the value is in the slice
	s4 = append(s4, 1000) // capacity is full, so s4 gets a new, bigger array; original is not changed

	fmt.Printf("s4 (modified original[:]): %v, length: %d, capacity: %d\n", s4, len(s4), cap(s4))
}
