package main

import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 19
	fmt.Println("First number is:", numbers)

	var names [5]string
	names[2] = "Alice"
	fmt.Println("Third name is:", names)

	var num = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array:", num)
	fmt.Println("Length of numbers array:", len(num))
	fmt.Println("Capacity of numbers array:", cap(num))
	fmt.Println("Type of numbers array:", num[2])

	var values [5]bool
	fmt.Println("Values array:", values)

	var fruits [5]string
	fmt.Println("Name of fruits :", fruits)
	fmt.Printf("Name of fruits %q:", fruits)
}
