package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user inpput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || err err

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating: ", input)
	fmt.Printf("type of this rating is %T:  ", input)

}
