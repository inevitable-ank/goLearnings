package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web request in go")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error occurred while making GET request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response Headers:", res.Header)
	fmt.Println("Response raw:", res)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(data))

}
