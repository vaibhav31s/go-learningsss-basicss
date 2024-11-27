package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Weclome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.November, 27, 23, 12, 10, 2, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006  Monday"))

}
