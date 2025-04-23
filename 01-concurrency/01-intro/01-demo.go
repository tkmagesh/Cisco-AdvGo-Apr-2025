package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of f1() through the scheduler to be executed in future
	f2()

	// blocking the execution of main() so that the scheduler can look for other goroutines (f1 in this case) that are scheduled and execute them
	// "poor man's synchronization technique" - DO NOT use this in production
	// time.Sleep(1 * time.Second)
	time.Sleep(4 * time.Second)
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
