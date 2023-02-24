package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const myURL = "http://localhost:8000/post"

func main() {
	fmt.Println("Welcome to web post course !")
	PerformPostJsonRequest(myURL)
}

func PerformPostJsonRequest(url string)  {
	// Let's create a fake payload here
	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with Golang",
			"price": 0,
			"platform": "learncodeonline.in"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, _ := io.ReadAll(response.Body)

	var responseString strings.Builder
	responseString.Write(data)

	fmt.Println(responseString.String())
}
