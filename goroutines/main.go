package main

import (
	"fmt"
	"time"
)

func sayHi() {
	fmt.Println("Hi there!")
	time.Sleep(1 * time.Second)
	fmt.Println("Timer for Hi ended")
}

func sayHello() {
	fmt.Println("Hello there!")
	time.Sleep(3 * time.Second)
	fmt.Println("Timer for Hello ended")
}

func main() {
	fmt.Println("Learning go routines")
	go sayHi()
	go sayHello()

	time.Sleep(2 * time.Second)

}