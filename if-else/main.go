package main

import "fmt"

func main() {
	num := 4
	if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}

	age := 20
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}

	if num >= 5 && age >= 18 {
		fmt.Println("Number is greater than or equal to 5 and age is 18 or older.")
	} else if num < 5 || age < 18 {
		fmt.Println("Either number is less than 5 or age is less than 18.")
	}
}
