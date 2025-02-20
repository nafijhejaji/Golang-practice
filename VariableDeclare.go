//here we will see how to declare variables in go

// so start with the basics
package main

import "fmt"

func main() {

	var a int = 20 //var=variable, a=variable name, int=type of variable, 20=value of variable

	// here we have declared a variable a of type int and assigned a value of 20 to it
	// we can also declare multiple variables in a single line

	var b, c int = 1, 2

	// here we have declared two variables b and c of type int and assigned values of 1 and 2 to them respectively

	// but go is very smart and can infer the type of variable by itself

	var d = 20

	// here we have declared a variable d and assigned a value of 20 to it
	// go will automatically infer the type of variable d as int

	// even if we don't declare anything though go will understand
	x := 100 // here we have declared a variable x and assigned a value of 100 to it
	x = 50   //we cant use (:) when we are assigning the value of a same variable the second time

	//but we must need to use the same data type of variable when we are assigning the value of a same variable the second time

	//here we didn't declare the type of variable x but go will automatically infer the type of variable x as int

	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d)
	fmt.Println(x)

}
