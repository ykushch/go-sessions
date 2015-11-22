package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

// defining new type with methods on it
type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

// defining interface with types
type geometry interface {
	calculateArea() float64
	calculatePerim() float64
}

type rect struct {
	width, height float64
}

type circ struct {
	radius float64
}

func (r rect) calculateArea() float64 {
	return r.width * r.height
}

func (r rect) calculatePerim() float64 {
	return 2*r.width + 2*r.height
}

func (c circ) calculateArea() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circ) calculatePerim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("Shape of type %s with geometry %v\n", reflect.TypeOf(g).String(), g)
	fmt.Println("Area is", g.calculateArea())
	fmt.Println("Perimeter is", g.calculatePerim())
}

// Working with errors
type argError struct {
	arg int
	err string
}

func f1(arg0 int) (int, error) {
	if arg0 == 42 {
		return -1, errors.New("Can't work with 42 input")
	}
	return arg0 + 1, nil
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.err)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 2, nil
}

func main() {
	fmt.Println("-> Working with structures")
	r := rectangle{width: 5, height: 6}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perimeter())

	r1 := &r
	fmt.Println("area: ", r1.area())
	fmt.Println("perim:", r1.perimeter())

	fmt.Println("-> Working with interfaces and structs")
	r2 := rect{width: 2, height: 4}
	c1 := circ{radius: 4}
	measure(r2)
	measure(c1)

	fmt.Println("-> Working with errors")
	fmt.Println(f1(42))
	// A nil value in the error position indicates that there was no error.
	fmt.Println(f1(41))

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.err)
	}

}
