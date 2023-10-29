package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to golang")

	// no inheritance in golang; no super or parent
	hitesh := User{"Hitesh", "raj@gmail.com", true, 28}

	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v", hitesh)

	hitesh.GetStatus()
	hitesh.NewEmail()

	fmt.Println(hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "fum@gmail.com"
	fmt.Println("Is User Active: ", u.Email)
}
