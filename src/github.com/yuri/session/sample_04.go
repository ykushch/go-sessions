package main

import "fmt"

func main() {
	// endless loop
	//for {
	//	fmt.Print("Hello-world\n")
	//}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	for i, j := 0, 1; i < 10; i, j = i + 1, j * 2 {
		fmt.Printf("Several indexes: %d - %d\n", i, j)
	}
}
