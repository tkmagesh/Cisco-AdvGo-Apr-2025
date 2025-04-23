/*
Follow "share memory by communicating" by removing the "primes[]" common variable and make this work
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 100)
	for no := range primes {
		fmt.Printf("Prime No : %d\n", no)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) chan int {
	primes := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go func(n int) {
				fmt.Printf("Goroutine started for [%d]\n", n)
				defer wg.Done()
				defer fmt.Printf("Goroutine completed for [%d]\n", n)
				if isPrime(n) {
					primes <- n
				}
			}(no)
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
