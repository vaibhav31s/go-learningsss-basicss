package main

import "fmt"

func main() {
	defer fmt.Println("Defer in Go 1")
	defer fmt.Println("Defer in Go 2")
	defer fmt.Println("Defer in Go 3")
	fmt.Println("Welcome to defer in Go")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Defer in Go", i)
	}
}
