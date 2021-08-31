package main

import (
	"fmt"
)

func main() {

	done := make(chan interface{})
	stringStream := make(chan string)

	go func() {
		stringStream <- "a"
		stringStream <- "a"
		close(done)
	}()

	for {
		select {
		case <-done:
			return
		case s := <-stringStream:
			fmt.Println(s)
		}
	}
}
