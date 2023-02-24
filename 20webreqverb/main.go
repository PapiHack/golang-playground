package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myGetURL = "http://localhost:8000/get"
const myPostURL = "http://localhost:8000/post"
const myPostFormURL = "http://localhost:8000/postform"

func main() {
	fmt.Println("Welcome to web verb video !")
	PerformGetRequest(myGetURL)
	fmt.Println("==========================")
	fmt.Println("Making a post request")
	PerformPostJsonRequest(myPostURL)
	fmt.Println("==========================")
	fmt.Println("Making a post form request")
	PerformPostFormRequest(myPostFormURL)
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

func PerformPostJsonRequest(url string)  {
	// Let's create a fake payload here
	requestPayload := strings.NewReader(`{
		"coursename": "Let's learn Golang",
		"price": 0,
		"platform": "learncodeonline"
	}`)

	response, err := http.Post(url, "application/json", requestPayload)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	responseByte, _ := io.ReadAll(response.Body)

	var responseString strings.Builder
	responseLength,_ := responseString.Write(responseByte)

	fmt.Println("Response Length is:", responseLength)
	fmt.Println(responseString.String())
	
}

func PerformPostFormRequest(hostURL string) {
	// Creating the FormData
	data := url.Values{}
	data.Add("firstname", "Elliot")
	data.Add("lastname", "Alderson")
	data.Add("email", "fsociety.dat")

	response, err := http.PostForm(hostURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println("Response is:", string(content))
}
