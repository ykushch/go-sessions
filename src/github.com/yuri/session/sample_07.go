package main

import (
	"fmt"
)

func main() {
	numbers := make(map[string]int)
	numbers["First"] = 1
	numbers["Second"] = 2
	numbers["Three"] = 3

	fmt.Printf("Second: %d\n", numbers["Second"])

	fmt.Printf("Second: %d\n", numbers["Five"])

	// distinguish
	val, ok := numbers["Five"]
	if !ok {
		fmt.Print("Can't get value")
	} else {
		fmt.Printf("%s", val)
	}

	// range
	has2 := 0

	delete(numbers, "Second")

	for _, val := range numbers {
		if val == 2 {
			has2++
		}
	}

	fmt.Printf("has2: %d\n", has2)

	nums := map[string]int {
		"First": 1,
		"Second": 2,
		"Three": 3,
	}

	fmt.Printf("First: %d\n", nums["First"])



}
