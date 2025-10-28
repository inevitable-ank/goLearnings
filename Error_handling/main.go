package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	// result, err := divide(10, 0)
	result, _ := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	fmt.Println("Result:", result)
	
}