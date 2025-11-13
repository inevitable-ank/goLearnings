package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getRequest() {
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

func postRequest() {
	// to be implemented
	todo := Todo{
		UserID:    1,
		Title:     "Learn Go Language",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)

	defer res.Body.Close()
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	fmt.Println("Response Status Code:", res.StatusCode)
	ioData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(ioData))
}

func main() {
	fmt.Println("Learning crud")
	// getRequest()
	postRequest()

}
