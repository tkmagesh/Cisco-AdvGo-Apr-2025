package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// non-blocking operation as the channel can receive it and hold
	ch <- 100
	data := <-ch
	fmt.Println(data)
}
