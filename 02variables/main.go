package main

import "fmt"

const LoginToken = "djflsiefj" // Public because of the first capital letter

func main() {
	var username string = "PapiHack"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.445545467521897
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default values & some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// Implicit type or type inference

	var webseite = "learngolang.dev"
	fmt.Println(webseite)

	// No var style (Not allowed outside of a function)

	numberOfuser := 30000
	fmt.Println(numberOfuser)
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}