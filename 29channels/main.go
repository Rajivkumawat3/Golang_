package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang --  lerncodeonline.in")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-mych

		fmt.Println(isChanelOpen)
		fmt.Println(val)
		// fmt.Println(<-mych)

		wg.Done()
	}(mych, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// mych <- 0
		close(mych)
		// mych <- 5
		// mych <- 3
		// close(mych)
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
