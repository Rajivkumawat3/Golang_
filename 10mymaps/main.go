package main

import (
	"fmt"
)

func main() {

	languages := make(map[string]string)

	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println(languages)

	delete(languages, "rb")

	fmt.Println(languages)

	// iterating in golang is really interesting

	for key, value := range languages {
		fmt.Printf("for key %v, the value is %v \n", key, value)
	}
}
