/*
modify the below to use context
*/
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type AppLogger struct {
	log *log.Logger
}

func NewAppLogger() *AppLogger {
	return &AppLogger{
		log: log.Default(),
	}
}

func (appLogger *AppLogger) Log(m ...any) {
	appLogger.log.Println(m...)
}

func main() {
	logger := NewAppLogger()
	/*

		logger.Log("Application started")
	*/
	/*
		rootCtx := context.Background()
		valCtx := context.WithValue(rootCtx, "root-key", "root-value")
		logCtx := context.WithValue(valCtx, "logger", logger)
	*/
	ctx := context.Background()
	ctx = context.WithValue(ctx, "root-key", "root-value")
	ctx = context.WithValue(ctx, "logger", logger)
	ctx = stop(ctx)
	ch := genNos(ctx)
	for data := range ch {
		fmt.Println(data)
	}
}

func stop(ctx context.Context) context.Context {
	logger := ctx.Value("logger").(*AppLogger)
	logger.Log("[stop] root-key :", ctx.Value("root-key"))
	cancelCtx, cancel := context.WithCancel(ctx)
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()
	return cancelCtx
}

func genNos(ctx context.Context) <-chan string {
	logger := ctx.Value("logger").(*AppLogger)
	logger.Log("[genNos] root-key :", ctx.Value("root-key"))
	genValCtx := context.WithValue(ctx, "root-key", "new-root-value")
	resultCh := make(chan string)
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		oddTimeoutCtx, cancel := context.WithTimeout(genValCtx, 10*time.Second)
		defer cancel()
		go genOddNos(oddTimeoutCtx, resultCh, wg)

		wg.Add(1)
		evenTimeoutCtx, cancel := context.WithTimeout(genValCtx, 15*time.Second)
		defer cancel()
		go genEvenNos(evenTimeoutCtx, resultCh, wg)

		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}

func genEvenNos(ctx context.Context, resultCh chan<- string, wg *sync.WaitGroup) {
	logger := ctx.Value("logger").(*AppLogger)
	logger.Log("[genEvenNos] root-key :", ctx.Value("root-key"))
	defer wg.Done()
LOOP:
	for no := 0; ; no += 2 {
		select {
		case resultCh <- fmt.Sprintf("Even : %d", no):
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			logger.Log("[genEvenNos] cancellation signal received")
			break LOOP
		}
	}
}

func genOddNos(ctx context.Context, resultCh chan<- string, wg *sync.WaitGroup) {
	logger := ctx.Value("logger").(*AppLogger)
	logger.Log("[genOddNos] root-key :", ctx.Value("root-key"))
	defer wg.Done()
LOOP:
	for no := 1; ; no += 2 {
		select {
		case resultCh <- fmt.Sprintf("Odd : %d", no):
			time.Sleep(500 * time.Millisecond)
		case <-ctx.Done():
			logger.Log("[genOddNos] cancellation signal received")
			break LOOP
		}
	}
}
