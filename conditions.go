package main

import "fmt"

func main() {
	age := 30

	if age > 18 {
		fmt.Println("you are adult")
	} else if age < 18 {
		fmt.Println("you are a kid")
	} else {
		fmt.Println("you are anything but not an adult or a kid")
	}
}
