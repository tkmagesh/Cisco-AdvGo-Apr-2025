/*
Follow "share memory by communicating" by removing the "primes[]" common variable and make this work
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 100, 5)
	for no := range primes {
		fmt.Printf("Prime No : %d\n", no)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int, workerCount int) chan int {
	primes := make(chan int)
	dataCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()
	go func() {
		wg := sync.WaitGroup{}
		for worker_id := range workerCount {
			wg.Add(1)
			go func() {
				fmt.Printf("Goroutine [id = %d] started\n", worker_id)
				defer wg.Done()
				defer fmt.Printf("Goroutine [id = %d] completed \n", worker_id)
				for no := range dataCh {
					fmt.Printf("Goroutine [id = %d] processing no : %d\n", worker_id, no)
					if isPrime(no) {
						primes <- no
					}
				}
			}()
		}
		wg.Wait()
		close(primes)
	}()
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
