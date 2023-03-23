package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)
	waitGroup := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	waitGroup.Add(2)
			// Receive only
	go func (ch <-chan int, wg *sync.WaitGroup)  {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		waitGroup.Done()
	}(myCh, waitGroup)
			// Send only
	go func (ch chan<- int, wg *sync.WaitGroup)  {
		myCh <- 5
		// myCh <- 6
		close(myCh)
		waitGroup.Done()
	}(myCh, waitGroup)

	waitGroup.Wait()
}