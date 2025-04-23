package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop!")
	go func() {
		fmt.Scanln()
		// close(stopCh)
		stopCh <- struct{}{}
	}()
	ch := genNos(stopCh)
	for no := range ch {
		fmt.Println("No :", no)
	}
}

func genNos(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-stopCh:
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- i * 10
			}
		}
		close(ch)
	}()
	return ch
}
