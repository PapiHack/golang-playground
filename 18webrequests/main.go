package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response if of type: %T\n", response)
	defer response.Body.Close()

	dataByte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataByte)

	fmt.Println(content)
}
