package main

import (
	"fmt"
	"reflect" // used to get the data type
)

// 4 ways to declare variables in go
var a = 5
var b int = 6

func main()  {
	c := 7 //this way only works within a function
	var d float64 //this way only works within a function
	d = 8.01

	// new examples within a function
	e := true
	f := "Hello World!"
	g := a + b
	h := string("33")

	fmt.Println(a, b, c, d, e, f, g, h, reflect.TypeOf(h))
}

// Printing outputs don't work outside a function.

// "fmt.Print" simply prints everything.
// "fmt.Println" adds a new line for each print.
