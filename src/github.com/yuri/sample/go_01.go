package main

import "fmt"
import "math"

func main() {
	var a int = 5
	var a_1 = 6
	fmt.Println("Simple variables addition: ", (a + a_1))

	var b string = "hello"
	var b_1 = "world"
	fmt.Println("Simple variables addition: ", (b + " " + b_1))

	// declaring multiple variables
	var multiple1, multiple2 int = 1, 2
	fmt.Println("The same but with multiple variables:", (multiple1 + multiple2))

	// the same as (will infer types automatically)
	// var shorthand string = "this works too"
	shorthand := "this works too"
	fmt.Println("Shorthand syntax: ", shorthand)

	const c1 = "I am constant"
	fmt.Println("Constant demo ", c1)

	const n = 500000
	const d = 3e20 / n
	fmt.Println("Operation with constants", d)

	// A numeric constant has no type until it’s given one,
	// such as by an explicit cast.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context
	// that requires one, such as a variable assignment or function call.
	// For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}
