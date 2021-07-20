package channel

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		n := <-ch1
		time.Sleep(time.Second)
		ch2 <- n * n
	}()
	ch1 <- 9

	assert.Equal(t, 81, <-ch2)
}

func TestChanClose(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for value := range ch1 {
			ch2 <- value * value
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("test")
			ch1 <- i
		}
		close(ch1)
	}()

	var ints []int
	for v := range ch2 {
		ints = append(ints, v)
	}

	assert.Contains(t, ints, 81)
	assert.NotContains(t, ints, 100)
}

func TestSelect(t *testing.T) {
	square := func(wg *sync.WaitGroup, ch chan int, quit chan bool) {
		for {
			select {
			case n := <-ch:
				fmt.Printf("Square: %d\n", n*n)
				time.Sleep(time.Second)
			case <-quit:
				wg.Done()
				return
			}
		}
	}

	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}

func TestSelectTick(t *testing.T) {
	square := func(wg *sync.WaitGroup, ch chan int) {
		tick := time.Tick(time.Second)
		terminate := time.After(10 * time.Second)

		for {
			select {
			case <-tick:
				fmt.Println("tick")
			case <-terminate:
				fmt.Println("terminated")
				wg.Done()
			case n := <-ch:
				fmt.Printf("square: %d\n", n*n)
				time.Sleep(time.Second)
			}
		}
	}

	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}

func TestProducerConsumer(t *testing.T) {
	type Car struct {
		Body    string
		Tire    string
		Color   string
		Success bool
	}
	wg := sync.WaitGroup{}
	startTime := time.Now()

	MakeBody := func(tireCh chan *Car) {
		tick := time.Tick(time.Second)
		after := time.After(10 * time.Second)
		for {
			select {
			case <-tick:
				car := &Car{}
				car.Body = "Sports car"
				tireCh <- car
			case <-after:
				close(tireCh)
				wg.Done()
				return
			}
		}
	}

	InstallTire := func(tireCh, paintCh chan *Car) {
		for car := range tireCh {
			time.Sleep(time.Second)
			car.Tire = "Winter tire"
			paintCh <- car
		}
		wg.Done()
		close(paintCh)
	}

	PaintCar := func(paintCh chan *Car) {
		for car := range paintCh {
			time.Sleep(time.Second)
			car.Color = "red"
			car.Success = true
			duration := time.Now().Sub(startTime)
			fmt.Printf("%.2f complete car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
		}
		wg.Done()
	}

	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("start factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}
