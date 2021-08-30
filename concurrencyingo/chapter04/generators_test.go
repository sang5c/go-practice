package chapter04

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRepeatAndTake(t *testing.T) {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	done := make(chan interface{})
	defer close(done)
	//values := []int{1, 2, 3, 4, 5}

	//go func() {
	//	time.Sleep(5 * time.Second)
	//	close(done)
	//}()
	//
	//result := repeat(done, values)
	//for v := range result {
	//	fmt.Println(v)
	//}

	// take
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	//done := make(chan interface{})
	//values := make(chan interface{}, 3)
	//values <- 1
	//values <- 2
	//values <- 3
	//
	//result := take(done, values, 2)
	//for v := range result {
	//	fmt.Println(v)
	//}

	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}

	toString := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	fmt.Println("\n---- toString ----")
	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}

func TestRepeatFn(t *testing.T) {
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} {
		return rand.Int()
	}

	for num := range take(done, repeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}
