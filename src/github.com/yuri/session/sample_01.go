package main

import "fmt"

func main() {
	atoz := "The quick brown fox jumps over the lazy dog\n"

	fmt.Print("Playing with Strings\n")

	fmt.Printf("%s\n", atoz[0:9])
	fmt.Printf("The same: %s\n", atoz[:9])

	fmt.Printf("Length: %d\n", len(atoz))

	for i, value := range atoz {
		fmt.Printf("%d - %c\n", i, value)
	}

	for _, value := range atoz {
		fmt.Printf("%c\n", value)
	}

	atoz = `The quick brown fox jumps over the lazy dog\n`
	fmt.Printf("With single qoute: %s\n", atoz)
}
