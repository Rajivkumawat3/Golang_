package main

import (
	"fmt"
	"net/url"
)

const myurls string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {

	fmt.Println("welcome to handling urls in golang: ")

	fmt.Println(myurls)

	// parsing

	result, _ := url.Parse(myurls)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("the type of qparams are: %T", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsofurls := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherURL := partsofurls.String()
	fmt.Println(anotherURL)
}
