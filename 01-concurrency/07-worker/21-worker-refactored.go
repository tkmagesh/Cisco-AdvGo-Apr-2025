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

func dataProducer(start, end int) <-chan int {
	dataCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()
	return dataCh
}

func genPrimes(start, end int, workerCount int) chan int {
	primes := make(chan int)
	dataCh := dataProducer(start, end)
	go func() {
		wg := &sync.WaitGroup{}
		for worker_id := range workerCount {
			wg.Add(1)
			go primeWorker(worker_id, dataCh, primes, wg)
		}
		wg.Wait()
		close(primes)
	}()
	return primes
}

func primeWorker(worker_id int, dataCh <-chan int, primes chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker [id = %d] started\n", worker_id)
	defer fmt.Printf("Worker [id = %d] completed \n", worker_id)
	for no := range dataCh {
		fmt.Printf("Worker [id = %d] processing no : %d\n", worker_id, no)
		if isPrime(no) {
			primes <- no
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
