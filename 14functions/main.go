package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is:", result)
	
	proResult, proMessage := proAdder(2, 5, 8, 7)
	fmt.Println("ProResult is:", proResult)
	fmt.Println("ProMessage is:", proMessage)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi, pro result function"
}

func greeter()  {
	fmt.Println("Namastey from Golang")
}
