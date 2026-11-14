package main

import "fmt"

func main() {

	// for is the only loop keyword in Go; it covers every kind of loop
	// C-style loop: init; condition; post
	for i := 0; i < 10; i++ { // i only exists inside the loop
		//fmt.Println(i) // commented out, so this loop prints nothing
	}

	// while-style: only a condition; Go has no while keyword
	k := 3
	for k > 0 {
		fmt.Println(k)
		k-- // decrease k by 1; without this the loop would never end
	}
	fmt.Println("-------Infinite loop-------")

	counter := 0
	for { // no condition: loops forever until break
		fmt.Println("counter's value is: ", counter)
		counter++
		if counter >=5 {
			break // exits the loop immediately
		}
	}

	fmt.Println("-------Skipping-------")

	for i := 0; i < 10; i++ {
		if i%2 == 0 { // % gives the remainder; even numbers have a remainder of 0
			continue // skips the rest of this round and goes to the next i
		}
		fmt.Println(i)
	}

	fmt.Println("-------Array-------")

	items := [3]string{"Go", "Python", "Java"} // array with a fixed length of 3
	for _, value := range items { // range gives the index and the value; _ ignores the index
		fmt.Println("Value: ", value)
	}

}
