package main

import (

)
import "fmt"

func emit(c chan string) {
	words := []string {"The", "black", "fox"}

	for _, word := range words {
		c <- word
	}

	close(c)
}

func main() {
	wordChan := make(chan string)
	// start go-routine
	go emit(wordChan)


	for i, word := range wordChan {
		if i == 2 {
			break
		}
		fmt.Printf("%s ", word)

	}

	word := <- wordChan

	fmt.Printf("%s\n", word)



}
