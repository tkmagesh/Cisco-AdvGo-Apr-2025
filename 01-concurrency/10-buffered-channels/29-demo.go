package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	// user is processing the error from the errCh
	/*
		oddNoCh, errCh := generateOdd()
			select {
			case oddNo := <-oddNoCh:
				fmt.Println("Odd number generated :", oddNo)
			case err := <-errCh:
				fmt.Println("error generating odd number :", err)
			}
	*/

	// User is not worried about the error
	oddNoCh, _ := generateOdd()
	/*
		select {
		case oddNo := <-oddNoCh:
			fmt.Println("Odd number generated :", oddNo)
		case <-time.After(2 * time.Second):
			fmt.Println("No odd number received")
		}
	*/
	fmt.Println("Odd number generated :", <-oddNoCh)
	fmt.Println("Done")
}

func generateOdd() (<-chan int, <-chan error) {
	errCh := make(chan error, 1)
	resultCh := make(chan int)
	go func() {
		if no := rand.Intn(20); no%2 == 0 {
			fmt.Println("[generateOdd] sending error")
			e := errors.New("error generating odd number")
			errCh <- e
		} else {
			resultCh <- no
		}
	}()
	return resultCh, errCh
}
