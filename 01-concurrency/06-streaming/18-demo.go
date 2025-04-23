package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := genNos()
	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Done")
}

func genNos() <-chan int {
	ch := make(chan int)
	// producer
	go func() {
		count := rand.Intn(20)
		fmt.Printf("[producer] producing %d numbers\n", count)
		for i := range count {
			ch <- (i + 1) * 10
		}
		fmt.Println("[producer] all the data have been produced, cloing the channel!")
		close(ch)
	}()
	return ch
}
