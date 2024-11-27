package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File handling in Go")
	content := "This needs to go in a file - Vaibhav.in"

	file, err := os.Create("./mylcogofile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length of content is: ", length)

	defer file.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	file, err := os.ReadFile(filename)

	checkNilErr(err)
	fmt.Println("Content of file is: ", string(file))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
