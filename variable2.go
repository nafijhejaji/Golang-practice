package main

import "fmt"

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	div := num1 / num2
	return sum, div
}

func main() {
	x := 30
	y := 10
	p, q := getNumbers(x, y)
	fmt.Println(p, q)

}
