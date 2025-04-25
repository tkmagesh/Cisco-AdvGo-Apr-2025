package main

import (
	"fmt"
	"testing-demo/utils"
	"time"
)

type SystemTimeProvider struct {
}

func (stp SystemTimeProvider) GetCurrent() time.Time {
	return time.Now()
}

func main() {
	stp := SystemTimeProvider{}
	greeter := utils.NewGreeter("Suresh", stp)
	fmt.Println(greeter.Greet())
}
