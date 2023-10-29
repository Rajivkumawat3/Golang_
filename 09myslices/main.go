package main

import (
	"fmt"
	"sort"
)


func main() {
	fmt.Println("welcome to video on slice")

	var fruitList = []string{"apple", "tomato", "chilli"}

	fmt.Printf("the type of fruitList: %T ", fruitList)

	fruitList = append(fruitList, "mango", "onion")
	fmt.Println("the type of fruitList ", fruitList)

	// fruitList = append(fruitList[:3])
	// fmt.Println("the type of fruitList ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("the type of fruitList ", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println("the type of fruitList ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 124
	highScores[2] = 125
	highScores[3] = 126
	//highScores[4]=127

	highScores = append(highScores, 444, 443)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	// how to remove a value from slice based index

	var courses = []string{"java", "python", "c++"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
