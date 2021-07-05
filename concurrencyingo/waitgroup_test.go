package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	result := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping")
		result++
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping")
		result++
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")

	assert.Equal(t, 2, result)
}

func TestForGo(t *testing.T) {
	result := 0

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		result++
		fmt.Printf("hello from %v!\n", id)
	}

	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()

	assert.Equal(t, numGreeters, result)
}
