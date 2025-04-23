/*
modify the below to take advantage the go concurrency model
*/
package main

import "fmt"

func main() {
	primes := genPrimes(2, 100)
	for _, no := range primes {
		fmt.Printf("Prime No : %d\n", no)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) []int {
	primes := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
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
