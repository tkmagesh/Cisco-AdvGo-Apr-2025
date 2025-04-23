/*
modify the below to take advantage the go concurrency model
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 100)
	for _, no := range primes {
		fmt.Printf("Prime No : %d\n", no)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) []int {
	primes := make([]int, 0)
	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func(n int) {
			fmt.Printf("Goroutine started for [%d]\n", n)
			defer wg.Done()
			defer fmt.Printf("Goroutine completed for [%d]\n", n)
			if isPrime(n) {
				mutex.Lock()
				{
					primes = append(primes, n)
				}
				mutex.Unlock()
			}
		}(no)
	}
	wg.Wait()
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
