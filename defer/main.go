package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("This is the first line")
	result := add(5, 7)
	defer fmt.Println("The result of addition is:", result)
	defer fmt.Println("This is the deferred line")
	fmt.Println("This is the last line")
}


// multiple defer goes o stack.. LIFO.. last in first out