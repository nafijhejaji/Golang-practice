package main

import "fmt"

func main() {
	age := 30
	sex := "male"
	sex = "female"
	if age == 20 && sex == "female" {
		fmt.Println("You can get married!")
	} else if age > 25 || sex == "male" {

		fmt.Println("your age is over")
	}
}
