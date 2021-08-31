package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	donePreventLeak()
	dontKnowStop()
	stopUseDone()
}

func donePreventLeak() {
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

func dontKnowStop() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			for {
				randStream <- rand.Int()
			}
		}()
		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints: ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
}

func stopUseDone() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			for {
				select {
				case <-done:
					return
				case randStream <- rand.Int():
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints: ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	time.Sleep(1 * time.Second)
}
