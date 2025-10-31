package main

import "fmt"

func main() {

	numbers := []int{10, 20, 30, 40, 50}
	numbers = append(numbers, 100, 90, 80, 70, 60)
	fmt.Println("Number : ", numbers)
	fmt.Printf("Type of numbers : %T\n" , numbers)
	fmt.Println("Length of numbers : ", len(numbers))
	fmt.Println("Capacity of numbers: ", cap(numbers))


	// empty slice

	name := []string{}
	fmt.Println("Name :", name)
	fmt.Printf("Name %q :", name)



	integers := make([]int, 5, 10)
	fmt.Println("Integers :", integers)
	fmt.Println("Lenght of integers :", len(integers))
	fmt.Println("Capacity of integers :", cap(integers))


	integers = append(integers, 10, 20, 30, 40)
	fmt.Println("Integers :", integers)
	fmt.Println("Lenght of integers :", len(integers))
	fmt.Println("Capacity of integers :", cap(integers))


	integers = append(integers, 50)
	fmt.Println("Integers :", integers)
	fmt.Println("Lenght of integers :", len(integers))
	fmt.Println("Capacity of integers :", cap(integers))

	integers = append(integers, 60)
	fmt.Println("Integers :", integers)
	fmt.Println("Lenght of integers :", len(integers))
	fmt.Println("Capacity of integers :", cap(integers))


	var values []bool
	fmt.Println("Values :", values)
	fmt.Printf("Values %T:", values)
}

