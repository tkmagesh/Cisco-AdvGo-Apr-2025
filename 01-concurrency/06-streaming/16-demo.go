package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := genNos()
	for range 5 {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	// producer
	go func() {
		for i := range 5 {
			time.Sleep(500 * time.Millisecond)
			ch <- (i + 1) * 10
		}

	}()
	return ch
}
