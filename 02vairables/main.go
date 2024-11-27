package main

import "fmt"

const LoginToken string = "asdasdasdasdasdasdasd"

func main() {
	var username string = "Vaibhav Gawad"
	fmt.Println("Hello, ", username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println("Is user logged in: ", isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println("smallValue is:  ", smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float64 = 255.4553123123
	fmt.Println("smallFloat is:  ", smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// defalult values and some aliases
	var anotherVarible int
	fmt.Println("anotherVarible is:  ", anotherVarible)
	fmt.Printf("Variable is of type: %T\n", anotherVarible)

	//implicit type
	var website = "gawad.in"
	fmt.Println("Website is:  ", website)
	fmt.Printf("Variable is of type: %T\n", website)

	//no var style

	numberOfUsers := 300000
	fmt.Println("Number of users:  ", numberOfUsers)
	fmt.Printf("Variable is of type: %T\n", numberOfUsers)

	fmt.Println(LoginToken)
}
