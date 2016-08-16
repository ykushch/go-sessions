package main

import (
	"fmt"
)

func printFunc(arr []string) {
	for _, value := range arr {
		fmt.Print(value)
	}
	fmt.Print("\n")
}

func main() {
	// words := [...]string {"Hello", "World", "!"}
	// printFunc(words)

	w := make([]string, 4)
	w[0] = "The"
	w[1] = "quick"
	w[2] = "brown"
	w[3] = "fox"
	printFunc(w)

	w2 := make([]string, 4)
	w2 = append(w2, "!", "hello", "world", "!")
	printFunc(w2)

	// len, cap
	newSlice := make([]string, 4)
	copy(newSlice, w)
	newSlice[2] = "bye"
	printFunc(newSlice)
}
