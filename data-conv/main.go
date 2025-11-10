package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 42
	fmt.Println("Value of num:", num)
	fmt.Printf("Type of num: %T\n", num)

	// num = num + 1.12 // This will cause a compile-time error
	var num1 float64 = float64(num)
	fmt.Println("Value of num after conversion to float64:", num1)
	fmt.Printf("Type of num1: %T\n", num1)

	num1 = num1 + 1.12

	fmt.Println("Value of num1 after addition:", num1)


	var num2 int = 33;

	number := int(num1) + num2;
	fmt.Println("Value of num3 (int(num1) + num2):", number)
	fmt.Printf("Type of num3: %T\n", number)

	data := strconv.Itoa(num2);
	fmt.Println("String representation of num2:", data)
	fmt.Printf("Type of data: %T\n", data)

	sum := data + " is the string representation of num2."
	fmt.Println(sum)

	num3, _ := strconv.Atoi(data);

	num3 = num3 + 100;
	fmt.Println("Integer value converted back from string:", num3)
	fmt.Printf("Type of num3: %T\n", num3)

	numString := "3.14"
	numFloat, _ := strconv.ParseFloat(numString, 64);

	fmt.Println("Float value converted from string:", numFloat)
	fmt.Printf("Type of numFloat: %T\n", numFloat)
}