package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)
	fmt.Printf("Time type %T\n", currentTime)

	formattedTime := currentTime.Format("02-01-2006 15:04:05 Monday")
	fmt.Println("Formatted Time:", formattedTime)

	layout_time := "02-01-2006"
	datestr := "2023-12-25"

	parsedTime, _ := time.Parse(layout_time, datestr)
	fmt.Println("Parsed Time:", parsedTime)


	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("New Date after adding 24 hours:", new_date)

	new_Formatted_date := new_date.Format("02-01-2006 15:04:05 Monday")
	fmt.Println("New Formatted Date:", new_Formatted_date)
}
