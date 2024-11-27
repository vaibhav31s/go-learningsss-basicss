package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	greeter()

	fmt.Println("Addition of numbers", adder(2, 3))

	result, msg := proAdder(3, 5, 123, 12, 31, 23, 123)

	fmt.Println("Addition of numbers result", result, msg)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	sum := 0
	for _, values := range values {
		sum += values
	}
	return sum, "Hi How are you!!"
}

func greeterTwo() {
	fmt.Println("Anotehr method namaste!!")
}
func greeter() {
	fmt.Println("Hello World, namaste!!")
}
