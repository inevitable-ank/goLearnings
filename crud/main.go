package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning crud")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response code:", res.StatusCode)
	}

	// Data, error := io.ReadAll(res.Body)
	// if error != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }

	// fmt.Println("Data is : ", string(Data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("completed Response:- ", todo.Completed)
	fmt.Println("Title Response:- ", todo.Title)
}
