package main

import "fmt"

func ChangeValueByaddress(ptr *int) {
	*ptr = 100
}

func main() {
	// var num int;
	num := 42
	fmt.Println("Value of num:", num)

	numPtr := &num
	fmt.Println("Address of num:", numPtr)
	fmt.Println("Value at address of num:", *numPtr)


	ChangeValueByaddress(numPtr)
	fmt.Println("Value of num after ChangeValueByaddress:", num)
}