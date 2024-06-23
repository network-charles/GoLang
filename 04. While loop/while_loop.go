package main

import "fmt"

// a simple while example
var spam = 0

func main() {
	for spam < 5 {
		fmt.Println(spam)
		spam++
	}
}