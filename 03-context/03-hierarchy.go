/*
modify the below to use context
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := stop()
	ch := genNos(ctx)
	for data := range ch {
		fmt.Println(data)
	}
}

func stop() context.Context {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()
	return ctx
}

func genNos(ctx context.Context) <-chan string {
	resultCh := make(chan string)
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		oddTimeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		go genOddNos(oddTimeoutCtx, resultCh, wg)

		wg.Add(1)
		evenTimeoutCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()
		go genEvenNos(evenTimeoutCtx, resultCh, wg)

		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}

func genEvenNos(ctx context.Context, resultCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := 0; ; no += 2 {
		select {
		case resultCh <- fmt.Sprintf("Even : %d", no):
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println("[genEvenNos] cancellation signal received")
			break LOOP
		}
	}
}

func genOddNos(ctx context.Context, resultCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := 1; ; no += 2 {
		select {
		case resultCh <- fmt.Sprintf("Odd : %d", no):
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println("[genOddNos] cancellation signal received")
			break LOOP
		}
	}
}
