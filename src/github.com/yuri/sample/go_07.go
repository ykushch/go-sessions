package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	fmt.Println("-> Working with goroutines")
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	fmt.Println("-> Working with channels")
	info := make(chan string)
	go func() {
		info <- "hello"
	}()
	msg := <-info
	fmt.Println("Message from channel is:", msg)

	messages := make(chan string, 2)
	messages <- "hello!"
	messages <- "world!"
	fmt.Println("First value:", <-messages)
	fmt.Println("Second value:", <-messages)

	fmt.Println("-> Working with channel synchronization")
	done := make(chan bool, 1)
	go worker(done)

	<-done

	fmt.Println("-> Working with channel directions")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message passed")
	pong(pings, pongs)
	fmt.Println("Message from pongs channel:", <-pongs)

	fmt.Println("-> Working with Select")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
