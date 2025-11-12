package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning go")

	myUrl := "https://www.youtube.com/watch?v=vu6ZQ-t1sUk&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=26"
	fmt.Println("URL is:", myUrl)
	fmt.Printf("Type of URL is: %T\n", myUrl)

	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Parsed URL is:", parsedUrl)
	fmt.Printf("Type of parsed URL is: %T\n", parsedUrl)

	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)

	//Modifying url

	parsedUrl.Scheme = "http"
	parsedUrl.Host = "google.com"	
	parsedUrl.Path = "/search"
	parsedUrl.RawQuery = "q=golang"

	fmt.Println("Modified URL is:", parsedUrl.String())
}
