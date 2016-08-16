package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Print("Playing with Strings\n")

	// Causes Error!
	//if numberBytes, err := fmt.Print("Playing with Strings\n"); err != nil {
	//	os.Exit(1)
	//
	//}
	//
	//fmt.Printf("Printed %d number bytes", numberBytes);

	if numberBytes, err := fmt.Print("Playing with Strings\n"); err != nil {
		os.Exit(1)

	} else {
		fmt.Printf("Printed %d number bytes", numberBytes);
	}

	fmt.Print("Other approach\n")

	var numberBytes int;
	var theError error;

	if numberBytes, theError = fmt.Print("Playing with Strings\n"); theError != nil {
		os.Exit(1)

	} else {
		fmt.Printf("Printed %d number bytes", numberBytes);
	}

}
