package main

import (
	"fmt"
)

func main() {

	fmt.Println("welcome to if-else statement")

	logincount := 293

	var result string

	if logincount < 30 {
		result = "regular count"
		fmt.Println(result)
	} else if logincount < 5 {
		result = "count"
		fmt.Println(result)
	} else {
		result := "maa chuda"
		fmt.Println(result)

	}

	fmt.Println("nice work")

	if num := 3; num < 9 {
		fmt.Println("kutta")
	} else {
		fmt.Println("kutti")
	}
}
