package main

import "fmt"


// Switch case example

func main() {
	Day := 3
	switch Day {
	case 1:
		fmt.Println("Monday");
	case 2:
		fmt.Println("Tuesday");
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}

	month := "March"
	switch month {
	case "January", "February", "March":
		fmt.Println("First Quarter")
	case "April", "May", "June":
		fmt.Println("Second Quarter")
	case "July", "August", "September":
		fmt.Println("Third Quarter")
	case "October", "November", "December":
		fmt.Println("Fourth Quarter")
	default:
		fmt.Println("Invalid Month")
	}

	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing Cold")
	case temperature >= 0 && temperature <= 15:
		fmt.Println("Cold")
	case temperature > 15 && temperature <= 25:
		fmt.Println("Warm")
	case temperature > 25 && temperature <= 35:
		fmt.Println("Hot")
	default:
		fmt.Println("Scorching Hot")
	}

}