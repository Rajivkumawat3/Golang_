package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("gand mara 1 ke sath")
		fallthrough
	case 2:
		fmt.Println("gand mara 2 ke sath")
		fallthrough
	case 3:
		fmt.Println("gand mara aur 3 maa chuda 1 ke sath")
	default:
		fmt.Println("muh mein le le mera")

	}
}
