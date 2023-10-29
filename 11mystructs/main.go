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
}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
