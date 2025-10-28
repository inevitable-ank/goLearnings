package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name ? ")
	var name string
	fmt.Scan(&name)

	fmt.Println("My name is :- ", name)

	fmt.Println("What is your full name ?")
	var fullName string
	reader := bufio.NewReader(os.Stdin)

	fullName, _ = reader.ReadString('\n');

	fmt.Println("My full name is :- ", fullName)

	// fmt.Scan reads only until the first whitespace (space, tab, or newline).
// So it reads "Ankit" into name,
// but leaves the rest " Kumar Ranjan\n" still sitting in the input buffer (stdin).

// But since the buffer still had " Kumar Ranjan\n",
// ReadString('\n') instantly reads that leftover text —
// it doesn’t wait for new user input!
}
