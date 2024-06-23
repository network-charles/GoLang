package main

import "fmt"

var name = "Mary"
var password = "swordfish"

func main() {
	if name == "Mary" {
		fmt.Print("Hello Mary")
		if password == "swordfish" {
			fmt.Print("\nAccess granted")
		} else {
			fmt.Print("Access denied")
		}
	}
}
