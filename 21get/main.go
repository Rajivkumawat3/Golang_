package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcom to get request in golang")
	//performgetrequest()
	//performPostJsonRequest()
	performFormRequest()
}

func performgetrequest() {
	const myurl = "https://lco.dev:3000/learn"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code: ", response.StatusCode)

	fmt.Println("content length is: ", response.ContentLength)

	// var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	// bytecount, _ := responseString.Write(content)

	// fmt.Println("bytecount is: ", bytecount)
	// fmt.Println(responseString)

	fmt.Println(content)
	fmt.Println(string(content))
}

func performPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fack json payload

	requestBody := strings.NewReader(`
	{
		"coursename":"let's go with golang",
		"price":115,
		"plateform":"learncodingonline"
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(content)
	fmt.Println(string(content))
}

func performFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}

	data.Add("firstName", "Rajiv")
	data.Add("lastName", "Kumawat")
	data.Add("E-mail", "Rajiv@gmail.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(content)
	fmt.Println(string(content))
}
