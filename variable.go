package main

import "fmt"

func summation1(type1 int, type2 int) {
	fmt.Println(type1 + type2)

}

func summation2(type1 float64, type2 float64) {
	fmt.Println(type1 + type2)

}

//in this summation function , if we want to use same datatype like float and float or int and int then we can use this function. But if we want to use int and float64 then we need to transform the int into float using floa64() function.

func main() {
	x := 15
	y := 5
	z := 20.5
	summation1(x, y)
	summation2(float64(x), z)

}
