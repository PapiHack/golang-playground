package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("Fruit list length is:", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegy list is:", vegList)
	fmt.Println("Vegy list length is:", len(vegList))
}
