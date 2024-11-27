package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))

	fruitList = append(fruitList, "Banana", "Mango")

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))

	fruitList = append(fruitList[1:])
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list length is: ", len(fruitList))

	highScore := make([]int, 3)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465

	highScore = append(highScore, 555, 666, 777)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	sort.Ints(highScore)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index

	var courses = []string{"react", "angular", "vue", "svelte"}

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
