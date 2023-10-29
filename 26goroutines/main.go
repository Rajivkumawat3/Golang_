package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usually pointers
var mut sync.Mutex    //// usually pointers
var signals = []string{"test"}

func main() {
	// go greeter("hello")
	// greeter("world")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}

	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(web)

	}

	wg.Wait()

	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint....")

	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
