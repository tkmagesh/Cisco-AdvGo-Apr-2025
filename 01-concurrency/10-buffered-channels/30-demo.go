/*
Shutdown on receiving kill signal
DO NOT accept user input in the genNos() function
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Process id :", os.Getpid())
	stopCh := stop()
	// stopCh := make(chan struct{})
	ch := genNos(stopCh)
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("Stop signal received... shutting down!")

}

func stop() <-chan struct{} {
	stopCh := make(chan struct{})
	go func() {
		osIntCh := make(chan os.Signal, 1)
		signal.Notify(osIntCh, syscall.SIGINT)
		<-osIntCh
		close(stopCh)
	}()
	return stopCh
}

func genNos(stopCh <-chan struct{}) <-chan string {
	resultCh := make(chan string)
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go genOddNos(stopCh, resultCh, wg)

		wg.Add(1)
		go genEvenNos(stopCh, resultCh, wg)

		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}

func genEvenNos(stopCh <-chan struct{}, resultCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := 0; ; no += 2 {
		select {
		case resultCh <- fmt.Sprintf("Even : %d", no):
			time.Sleep(500 * time.Millisecond)
		case <-stopCh:
			fmt.Println("[genEvenNos] stop signal received.. shutting down!")
			break LOOP
		}
	}
}

func genOddNos(stopCh <-chan struct{}, resultCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := 1; ; no += 2 {
		select {
		case resultCh <- fmt.Sprintf("Odd : %d", no):
			time.Sleep(300 * time.Millisecond)
		case <-stopCh:
			fmt.Println("[genOddNos] stop signal received.. shutting down!")
			break LOOP
		}
	}
}
