package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	b :=make([]byte, 100)
	n, err := f.Read(b)

	fmt.Printf("%d: % x\n\n", n, b)

	stringVersion := string(b)
	fmt.Printf("%d: %s\n", n, stringVersion)

	someString := "foo bar"
	f.Write([]byte(someString))
}
