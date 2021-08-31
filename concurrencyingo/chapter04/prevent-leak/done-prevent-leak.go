package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWOrk exited.")
			defer close(terminated)
			for {
				select {
				case <-done:
					return
				case s := <-strings:
					fmt.Println(s)
				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})
	//strings := make(chan string)
	//terminated := doWork(done, strings)
	terminated := doWork(done, nil)

	//go func() {
	//	for i := 0; i < 100; i++ {
	//		strings <- time.Now().String()
	//		time.Sleep(200 * time.Millisecond)
	//	}
	//}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated // join
	fmt.Println("Done.")
}
