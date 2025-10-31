package main

import "fmt"

func main() {
	for i:= 1; i <= 5; i++ {
		fmt.Println(i);
	}

	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	data := "Hello, Go!"
	data1 := "I love go lang"

	for index, value := range data {
		fmt.Printf("index is : %d char is : %c\n", index, value)
	}

	for index, value := range data1 {
		fmt.Printf("index is : %d char is : %c\n", index, value)
	}
}