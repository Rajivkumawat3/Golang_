package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to defer course1")
	defer fmt.Println("welcome to defer course 2")
	defer fmt.Println("welcome to defer course 3")
	fmt.Println("welcome to defer course 4")

	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
