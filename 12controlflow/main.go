package main

import "fmt"

func main() {
	fmt.Println("Welcome to control flow in golang")

	loginCount := 10

	if loginCount <= 10 {
		fmt.Println("Regular User")
	} else {
		fmt.Println("Special User")
	}

	if loginCount%2 == 1 {
		fmt.Println("Odd User")
	} else {
		fmt.Println("Even User")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

}
