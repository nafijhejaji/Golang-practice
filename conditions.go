package main

import "fmt"

func main() {

	age := 20

	if age > 18 {
		fmt.Println("You are ready to go!")
	} else if age == 18 {
		fmt.Println("You are almost ready to go!")
	} else {
		fmt.Println("you are a child!")
	}

}
