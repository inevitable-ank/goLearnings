package main

import "fmt"

func Add(a, b int) (int) {
	return a + b
}

func Multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("Learing Golang")

	sum := Add(5, 7)
	fmt.Println("Sum is:", sum)


	product := Multiply(4, 6)
	fmt.Println("Product is:", product)
}