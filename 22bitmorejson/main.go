package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string   `json:"coursename"`
	Price     int      `json:"price"`
	Plateform string   `json:"plateform"`
	Password  string   `json:"-"`
	Tags      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to golang, how to create json data in golang")
	EncodeJson()
	//decodeJson()
}

func EncodeJson() {
	lcocourses := []course{
		{"reactjs", 234, "lcoyoutube", "123456", []string{"web-dev", "js"}},
		{"nodejs", 44, "lcoyoutube", "123456", []string{"web-dev", "js"}},
		{"redux", 123, "lcoyoutube", "123456", nil},
	}

	// package this data as json data

	fmt.Println(lcocourses)

	finaljson, _ := json.MarshalIndent(lcocourses, "", "\t")

	// because finaljson is in bytes

	fmt.Println(string(finaljson))

}

func decodeJson() {

	// the json data is comming from byte formate
	jsonDataFromWeb := []byte(`{
		"coursename": "reactjs",
		"price": 234,
		"plateform": "lcoyoutube",
		"tags": ["web-dev","js"]
         }
   `)
	fmt.Println(jsonDataFromWeb)
	var lcoCourse course // type of struct variable

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("json was valid ")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("json is not valid ")
	}

	// some cases where you just want to data to key value

}

// type Link struct { A, Z string }
// l := Link{A: "n1:2", Z: "n2:1"}
// fmt.Printf("%v\n", l)	// {n1:2 n2:1}
// fmt.Printf("%+v\n", l)	// {A:n1:2 Z:n2:1}
// fmt.Printf("%#v\n", l)	// pkg.Link{A:"n1:2", Z:"n2:1"}
