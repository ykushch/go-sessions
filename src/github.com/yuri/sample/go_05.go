package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
}

func stdFunc(val int) {
	val = 0
}

func ptrFunc(val *int) {
	*val = 0
}

func main() {
	i := 1
	fmt.Println("-> Demo with pointers")
	fmt.Println("Inital value", i)
	// pass by value
	stdFunc(i)
	fmt.Println("After applying stdFunc(i) value is", i)

	ptrFunc(&i)
	fmt.Println("After applying ptrFunc(&i) value is", i)

	fmt.Println("-> Working with structs")
	fmt.Println(person{name: "Yuri", surname: "Kushch"})
	fmt.Println(person{surname: "Kushch"})
	fmt.Println(&person{name: "Yuri", surname: "Kushch"})

	personObj := person{name: "Yuri", surname: "Kushch"}
	personPtr := &personObj
	personPtr.name = "IIURI"
	fmt.Println("personPtr", personPtr)
	fmt.Println("personObj", personObj)
}
