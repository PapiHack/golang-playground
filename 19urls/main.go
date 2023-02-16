package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=slaskalk"

func main() {
	fmt.Println("Welcome to handling urls in Golang!")
	fmt.Println(myURL)

	result, err := url.Parse(myURL)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qParams)

	fmt.Println(qParams["coursename"])

	for key, val := range qParams {
		fmt.Printf("The query parameter '%v' has the value '%v'.\n", key, val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "tutcss",
		RawPath: "user=papihack",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
