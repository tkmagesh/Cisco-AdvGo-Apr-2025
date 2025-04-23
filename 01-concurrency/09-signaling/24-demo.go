package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for no := range ch {
		fmt.Println("No :", no)
	}
}

func genNos() <-chan int {
	ch := make(chan int)
	// timeoutCh := timeout(5 * time.Second)
	timeoutCh := time.After(5 * time.Second)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-timeoutCh:
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

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
