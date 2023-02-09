package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang!")

	currentTime := time.Now()
	fmt.Println(currentTime)

	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.October, 16, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))
}