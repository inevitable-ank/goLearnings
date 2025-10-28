package main

import "fmt"

func main() {
	name := "Alice"
	age := 30
	height := 5.7123456789

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)

	fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)

	fmt.Printf("My Name is: %s", name)
	fmt.Printf(", I am %d years old", age)
	fmt.Printf(", and my height is %.20f feet.\n", height)
}