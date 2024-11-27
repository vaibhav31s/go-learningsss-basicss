package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	vaibhav := User{"Vaibhav", "gawadvaibhavv@gmail.com", true, 22}
	fmt.Println(vaibhav)
	fmt.Println("Name is: ", vaibhav.Name)
	vaibhav.NewEmail()
	fmt.Println("Email is: ", vaibhav.Email)
	fmt.Println("Status is: ", vaibhav.Status)
	fmt.Println("Age is: ", vaibhav.Age)

	fmt.Printf("Name is: %v, Email is: %v, Status is: %v, Age is: %v \n", vaibhav.Name, vaibhav.Email, vaibhav.Status, vaibhav.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) GetName() {
	fmt.Println("User name is: ", u.Name)
}

func (u User) NewEmail() {
	u.Email = "gasdfjas@gmail.coM"
	fmt.Println("New email is: ", u.Email)
}
