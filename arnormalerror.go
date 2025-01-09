package main

import "fmt"

func main3() {
	a := 284
	a = 20

	fmt.Print(a)

	a = 100
	// Here a = 100 will not be printed because a = 100 has been assigned after the first assignment of a = 20.
}
