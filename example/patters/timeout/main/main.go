package main

import (
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/deadline"
)

func main() {
	makeTimeoutRequest()
}

func makeTimeoutRequest() {
	dl := deadline.New(1 * time.Second)
	err := dl.Run(func(stopper <-chan struct{}) error {
		slowFunction()

		return nil
	})

	switch err {
	case deadline.ErrTimedOut:
		fmt.Println("TimeOut")
	default:
		fmt.Println(err)
	}
}

func slowFunction() {
	time.Sleep(100 * time.Second)
	return 
}