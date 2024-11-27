package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("Languages: ", languages)

	delete(languages, "RB")

	fmt.Println("JS is: ", languages["JS"])

	fmt.Println("Languages: ", languages)

	fmt.Println("Languages: ", languages)

	fmt.Println("Languages: ", languages["RB"])

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Println("For Key", key, "value is ", value)
	}
}
