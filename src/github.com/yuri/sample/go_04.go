package main

import (
	"fmt"
)

// show the basic usage of functions
// the semanthic is simple
// name(paramName paramType, paramName paramType) paramResultType {}
func plus(a int, b int) int {
	return a + b
}

// if all variable are of the same type
// then the type can safely be omitted like this
func plusWithImplcitFirstType(a, b, c int) int {
	return a + b + c
}

func main() {
	// function usage
	fmt.Println("-> Basic function usage")
	res := plus(5, 5)
	fmt.Println(res)

	res = plusWithImplcitFirstType(1, 2, 3)
	fmt.Println(res)

	fmt.Println("-> Function with multiple return types")
	res1, res2 := returnMultipleValues()
	fmt.Printf("Multiple return values are: %v and %v\n", res1, res2)
	// you can omit some values from functions
	// like you've seen before
	_, res3 := returnMultipleValues()
	fmt.Println("Value with omitted first return argument is", res3)

	fmt.Println("-> Variadic functions")
	totalVariadic := sumVariadic(1, 2, 3, 4, 5)
	fmt.Println("Variadic total is", totalVariadic)

	nums := []int{1, 2, 3}
	totalVariadic = sumVariadic(nums...)
	fmt.Println("Using slices:", totalVariadic)

	fmt.Println("-> Anonymous functions")
	nextInt := initSeq()
	fmt.Println("Value from anonymous function:", nextInt)
	fmt.Println("Next value is:", nextInt())
	fmt.Println("And the next one:", nextInt())

	fmt.Println("Restarting sequence")
	nextInt = initSeq()
	fmt.Println("First value is:", nextInt)
	fmt.Println("Next value is:", nextInt())

	fmt.Println("-> Recursion functions")
	fmt.Println("Recursive factorial call", recursiveFactorialCall(4))
}

// like in other languages function could be used
// after main function
func returnMultipleValues() (int, int) {
	return 5, 6
}

func sumVariadic(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func recursiveFactorialCall(number int) int {
	if number == 0 {
		return 1
	}
	return number * recursiveFactorialCall(number-1)
}
