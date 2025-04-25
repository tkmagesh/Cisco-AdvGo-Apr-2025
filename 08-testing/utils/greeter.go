package utils

import (
	"fmt"
)

type Greeter struct {
	UserName     string
	timeProvider TimeProvider
}

func (greeter *Greeter) Greet() string {
	if greeter.timeProvider.GetCurrent().Hour() < 12 {
		return fmt.Sprintf("Hi %s, Good Morning!", greeter.UserName)
	}
	return fmt.Sprintf("Hi %s, Good Day!", greeter.UserName)
}

func NewGreeter(userName string, timeProvider TimeProvider) *Greeter {
	return &Greeter{
		UserName:     userName,
		timeProvider: timeProvider,
	}
}
