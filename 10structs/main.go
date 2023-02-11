package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	papihack := User{"papihack", "papihack@fosciety.dat", true, 16}
	fmt.Println(papihack)
	fmt.Printf("papihack details info are: %+v\n", papihack)
	fmt.Printf("Name is \"%v\" and Email is \"%v\"\n", papihack.Name, papihack.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}
