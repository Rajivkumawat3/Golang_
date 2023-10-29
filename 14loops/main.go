package main

import (
	"fmt"
)

func main() {

	fmt.Println("welcome to loops in golang")

	//days := []string{"monday", "tuesday", "friday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index, day)
	// }

	rougevalue := 5

	for rougevalue < 10 {
		// if rougevalue==5{
		// 	break
		// }
		if rougevalue == 2 {
			goto lco
		}
		if rougevalue == 5 {
			rougevalue++
			continue
		}
		fmt.Println(rougevalue)
		rougevalue++
	}

lco:
	fmt.Println("jumping like monkey")
}
