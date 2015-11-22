package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Endless loops with break")
		break
	}

	// a statement can preceed conditionals
	// variable that are declered in first if
	// are available in all branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// no ternary (?:) operator

	i = 2
	// Switch
	fmt.Println("Switch statements")
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	}

	// case values could be comma-separated
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	// switch without an expression is
	// like the 'if/else' logic
	// case values could be non-constants
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

}
