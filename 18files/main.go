package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("welcome to files")

	content := "this code needs to go in files- learncodeonline@.in"

	file, err := os.Create("./mylcofiles.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)

	defer file.Close()

	readFile("./mylcofiles.txt")
}

func readFile(fileName string) {
	databytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))

}
