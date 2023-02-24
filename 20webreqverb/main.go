package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const myURL = "http://localhost:8000/get"

func main() {
	fmt.Println("Welcome to web verb video !")
	PerformGetRequest(myURL)
}

func PerformGetRequest(url string)  {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}
