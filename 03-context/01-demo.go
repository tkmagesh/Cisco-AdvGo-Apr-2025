package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to stop!")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	ch := genNos(cancelCtx)
	for no := range ch {
		fmt.Println("No :", no)
	}
}

func genNos(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("[genNos] cancellation signal received")
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
