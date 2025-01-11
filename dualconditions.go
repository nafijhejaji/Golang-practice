package main

import "fmt"

func main() {
	age := 20
	sex := "male"
	sex2 := "female"
	if age == 20 && sex == "male" {
		fmt.Println("You can get married!")
	} else if age == 20 || sex2 == "female" {

		fmt.Println("your age is over")
	}
}
