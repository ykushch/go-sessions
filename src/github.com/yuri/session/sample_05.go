package main

import (
	"fmt"
)

func printer(msg string) {
	fmt.Printf(msg)
}

func printer2(msg, msg2 string, repeat int) {
	for repeat > 0 {
		fmt.Printf("%s: %s\n", msg, msg2)
		repeat -= 1
	}

}

func printer3(msg string) (string, error) {
	msg += "\n"

	_, err := fmt.Printf("%s\n", msg)

	return msg, err
}

func printer4(msg string) error {
	defer fmt.Print("Hey, I'm the last one\n")

	_, err := fmt.Printf("%s\n", msg)

	return err
}

func printer5(msg string) (str string, e error) {
	_, e = fmt.Printf("%s\n", msg)
	return
}

func printer6(msg ...string) {
	for _, val := range msg {
		fmt.Printf("%s \n", val)
	}
}

func main() {
	fmt.Print("Hello world")

	// simple function call
	str := "123"
	printer(str)

	printer2(str, "321", 5)

	msg, err := printer3(str)
	if err == nil {
		fmt.Printf("%q", msg)
	}

	printer4("hello")

	printer5("Hello\n\n")

	printer6("Hello", "World", "!")
}
