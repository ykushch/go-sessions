package main

import (
	"os"
	"fmt"
	"errors"
)

var (
	errorEmptyString = errors.New("Not working with empty string")
)

func printErr(msg string) error {
	if msg == "" {
		// return fmt.Errorf("Not working with empty string")
		// panic(errorEmptyString)
		return errorEmptyString
	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printErr(""); err != nil {
		if err == errorEmptyString {
			fmt.Print("Empty String was occured")
		} else {
			fmt.Printf("Failed with error: %s\n", err)
		}

		os.Exit(1)
	}



}
