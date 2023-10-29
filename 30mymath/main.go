package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("welcome to mymaths in golang")

	// var myNumber int = 34
	// var myNumberr float64 = 2.4

	// fmt.Println(myNumber + int(myNumberr))

	/// random number

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) // range between 0+1----4+1

	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandNum)

}
