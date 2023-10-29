package main

import (
	"fmt"
)

func main() {
	greeter()

	result := adder(3, 5)

	fmt.Println(result)

	proRes, message := proresult(200, 300, 400, 500)
	fmt.Println(proRes)
	fmt.Println(message)
}

func greeter() {
	fmt.Println("welcome rajiv ")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proresult(value ...int) (int, string) {

	total := 0

	for _, val := range value {
		total += val
	}
	return total, "nice dude"
}
