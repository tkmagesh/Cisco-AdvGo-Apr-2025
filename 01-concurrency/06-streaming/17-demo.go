package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := genNos()
	for {
		time.Sleep(500 * time.Millisecond)
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			continue
		}
		break
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
