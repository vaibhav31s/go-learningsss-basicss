package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println("Using for loop")

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Println(index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 5 {
			goto outer
		}
		fmt.Println("values is : ", rougueValue)
		rougueValue++
	}

outer:
	fmt.Println("Outer loop  Executed!!")
}
