package main

import (
	"fmt"
	myutils "golang/myUtils"
)

func main() {
	fmt.Println("Hello, World!")

	myutils.Add(3, 4)

	var money int = 1000
	fmt.Println("Money:", money)

	var name string = "John"
	fmt.Println("Name:", name)

	var isActive bool = true
	fmt.Println("Is Active:", isActive)

	isActive = false

	const pi float64 = 3.14
	fmt.Println("Pi:", pi)

	const piApprox = 22.0 / 7.0
	fmt.Println("Pi Approximation:", piApprox);

	welcomeMessage := "Welcome to Go!"
	fmt.Println(welcomeMessage);
}
