package main

import "fmt"

func main() {
	type Student struct {
		Name  string
		Age   int
		Grade string
	}

	var student1 Student
	fmt.Println("Student1 :", student1)

	student1.Name = "Alice"
	student1.Age = 20
	student1.Grade = "A"
	fmt.Println("Student1 after initialization:", student1)

	student2 := Student{
		Name:  "Bob",
		Age:   22,
		Grade: "B",
	}
	fmt.Println("Student2 :", student2)

	var student3 = Student{"Charlie", 19, "A+"}
	fmt.Println("Student3 :", student3)

	var student4 = new(Student)
	student4.Name = "Diana"
	student4.Age = 21
	student4.Grade = "A-"
	fmt.Println("Student4 :", *student4)

	type ContactInfo struct {
		Phone   int
		Email   string
		ZipCode int
	}

	type Address struct {
		HouseNo int
		Street  string
		City    string
		State   string
		Country string
	}

	type Employee struct {
		Contact ContactInfo
		Address Address
		Name    string
		Age     int
		Position string
	}

	fmt.Println("Employee Struct Example:", Employee{})
	Employee1 := Employee{
		Contact: ContactInfo{
			Phone:   1234567890,
			Email:   "alpha@example.com",
			ZipCode: 12345,
		},
		Address: Address{
			HouseNo: 101,
			Street:  "Main St",
			City:    "Metropolis",
			State:   "StateX",
			Country: "CountryY",
		},
		Name:     "Alpha",
		Age:      30,
		Position: "Engineer",
	}
	fmt.Println("Employee1 Details:", Employee1)
}
