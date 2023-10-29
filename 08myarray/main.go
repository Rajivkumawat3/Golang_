package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to array")

	var fruitless [4]string

	fruitless[0] = "apple"
	fruitless[1] = "mango"
	fruitless[3] = "shake"

	fmt.Println("fruitlist is: ", fruitless)
	fmt.Println("fruitlist is: ", len(fruitless))

	var veglist = [3]string{"potato", "onion", "beans"}

	fmt.Println("fruitlist is: ", veglist)
	fmt.Println("fruitlist is: ", len(veglist))

}
