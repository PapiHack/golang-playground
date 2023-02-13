package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in Golang")
	papihack := User{"PapiHack", "papihack@fsociety.dat", true, 20}
	fmt.Printf("papihack details are: %+v\n", papihack)
	fmt.Printf("Name is %v and email is %v\n", papihack.Name, papihack.Email)
	papihack.GetStatus()
	papihack.NewMail()
	fmt.Printf("Name is %v and email is %v\n", papihack.Name, papihack.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus()  {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@fsociety.dev"
	fmt.Println("Email of this user is:", u.Email)
}
