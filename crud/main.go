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
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	defer res.Body.Close()
	
	fmt.Println("Response Status Code:", res.StatusCode)
	ioData, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(ioData))
}

func putRequest() {
	// to be implemented
	todo := Todo{
		UserID:    11234567890,
		Title:     "Learn Go Language - Updated",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status Code:", res.StatusCode)
	ioData, _ := io.ReadAll(res.Body)	
	fmt.Println("Response Body:", string(ioData))


}

func deleteRequest() {
	// to be implemented

	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)	
	if err != nil {
		fmt.Println("Error making DELETE request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status Code:", res.StatusCode)
	ioData, _ := io.ReadAll(res.Body)	
	fmt.Println("Response Body:", string(ioData))
}

func main() {
	fmt.Println("Learning crud")
	// getRequest()
	// postRequest()
	// putRequest()
	deleteRequest()

}
