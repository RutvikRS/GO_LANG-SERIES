package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go Lang")

	myCh := make(chan int, 5)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	// R ONLY
	go func(ch chan int, wg *sync.WaitGroup) {

		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 5
		// myCh <- 6

		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
