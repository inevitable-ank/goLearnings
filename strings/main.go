package main

import (
	"fmt"
	"strings"
)

func main() {
	fruits := "apple,banana,cherry"

	fmt.Println("Fruits:", fruits)
	splittedFruits := strings.Split(fruits, ",");

	fmt.Println("Splitted Fruits:", splittedFruits)


	str := " one two three four five two two two one three two"

	count := strings.Count(str, "two")
	fmt.Println("Count of 'two':", count)


	str = "                                   hello go lang                                   "
	fmt.Println("Trimmed text:- ", strings.TrimSpace(str))

	firstName := "John"
	LastName := "Doe"

	FullName := strings.Join([]string{firstName, LastName}, " ")
	fmt.Println("Full Name:", FullName)

}