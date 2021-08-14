package chapter03

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChan(t *testing.T) {
	stringStream := make(chan string)
	str := "Hello Channels!"
	go func() {
		stringStream <- str
	}()

	value, _ := <-stringStream
	assert.Equal(t, str, value)
}

func TestCloseChannel(t *testing.T) {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, integer)
}

func TestChanRange(t *testing.T) {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 5; i++ {
			intStream <- i
		}
	}()

	for i := range intStream {
		fmt.Println(i)
	}
}
