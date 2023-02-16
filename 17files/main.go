package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang!")
	content := "This, needs to go in a file - fsociety.dat"

	file, err := os.Create("./file.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./file.txt")
}

func readFile(filename string) {
	byteDate, err := os.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is: \n", string(byteDate))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
