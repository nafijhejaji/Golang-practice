package main

import "fmt"

func main() {
	a := 13
	a = 100 // we cannot use (:) the second time we want to assign a value to a same variable.
	// Line 7 will print 100 not 13.
	fmt.Println(a) //the basic one

	var c int = 15
	var d = 40.56      // we can also use this way to declare a variable. Go doesn't require you to declare the type of the variable, it can infer it from the value you assign to it.
	e := "hello world" // we can also use this way to declare a variable. Go doesn't require you to declare the type of the variable, it can infer it from the value you assign to it.
	g := true          // the first time what kind of value you assign to the variable, the type of the variable will be the same as the value you assign to it. In this case, the type of g is bool. then you can't assign a string to g. g = "hello world" will cause an error. because the type of g is bool not string.
	g = false
	fmt.Println(c, d, e) // This will assign the values to the variables and print them out on the same line.
	fmt.Println(g)       // This will print the value of g which is false. Not true.

	const p = 100 // This is a constant variable. You can't change the value of a constant variable. If you try to change the value of a constant variable, you will get an error.
}
