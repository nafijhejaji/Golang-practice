package main

import "fmt"

func main() {

	a := 20
	switch a {
	case 1:
		fmt.Println("One")
	case 2, 3:
		fmt.Println("Two")
	default:
		fmt.Println("Neither case 1 nor case 2")
	}
}
