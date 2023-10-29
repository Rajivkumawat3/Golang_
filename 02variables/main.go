package main

import "fmt"

func main() {

	var username string = "rajiv kumawat"
	fmt.Println(username)
	fmt.Printf("variable type is: %T \n", username)

	var isloggedin bool = false
	fmt.Println(isloggedin)
	fmt.Printf("variable type is: %T \n", isloggedin)

	var number uint8 = 123
	fmt.Println(number)
	fmt.Printf("variable type is: %T \n", number)

	// default values and some aliases
	var intnumber string
	fmt.Println(intnumber)
	fmt.Printf("variable type is: %T \n", intnumber)

	//implicit declaration
	var payment = 3000
	fmt.Println(payment)
	fmt.Printf("variable type is: %T \n", payment)

	// no var style
	numberofusers := 3000
	fmt.Println(numberofusers)
	fmt.Printf("variable type is: %T \n", numberofusers)

}
