package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("Example.txt")
	if err != nil {
		fmt.Println("Error Creating file", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully:", file.Name())

	_, err = file.WriteString("Hello, this is an example of file handling in Go.\n")
	if err != nil {
		fmt.Println("Error writing to file", err)
		return
	}

	content := "This file demonstrates creating, writing, and closing a file using Go."

	bytes, err := io.WriteString(file, content)

	fmt.Println("Number of bytes written:", bytes)
	if err != nil {
		fmt.Println("Error writing to file", err)
		return
	}

	//// file read operations
	fileRead, err := os.Open("Example.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer fileRead.Close()

	buffer := make([]byte, 100)

	for {
		n, err := fileRead.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file", err)
			return
		}

		fmt.Print(string(buffer[:n]))
	}

	contentNew, err := ioutil.ReadFile("Example.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println("\nContent read using ioutil.ReadFile:")
	fmt.Println(string(contentNew))

}
