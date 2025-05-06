package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")
	mych := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// mych <- 5
	// fmt.Println(<-mych)
	wg.Add(2)
	// Recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(mych)
		val, isopen := <-mych
		fmt.Println(isopen)
		fmt.Println(val)
		// fmt.Println(<-mych)
		// fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)
	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 0
		close(mych)
		// mych <- 6
		wg.Done()
	}(mych, wg)
	wg.Wait()

}
