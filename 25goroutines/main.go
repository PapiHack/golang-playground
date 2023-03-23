package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // Usually pointer

func main() {
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, site := range websiteList {
		go getStatusCode(site)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint !")
		panic(err)
	}

	fmt.Printf("%d status code for %s !\n", response.StatusCode, endpoint)
}
