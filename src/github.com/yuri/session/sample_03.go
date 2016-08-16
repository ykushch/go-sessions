package main

import (
	"fmt"
	"time"
)

func main() {
	// no ternary (?:) operator

	i := 2
	// Switch
	fmt.Println("Switch statements")
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough
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