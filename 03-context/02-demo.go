package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create a root context
	rootCtx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()

	fmt.Println("Hit ENTER to stop (auto stops after 5 seconds)...")
	go func() {
		fmt.Scanln()
		cancel() // send the cancellation signal through the timeoutCtx
	}()

	ch := genNos(timeoutCtx)
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
				if ctx.Err() == context.Canceled {
					fmt.Println("[genNos] : programmatic cancellation initiated")
				}
				if ctx.Err() == context.DeadlineExceeded {
					fmt.Println("[genNos] : timeout based cancellation initiated")
				}
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
