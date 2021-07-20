package channel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	PrintEverySecond := func(ctx context.Context, wg *sync.WaitGroup) {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case <-tick:
				fmt.Println("tick")
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go PrintEverySecond(ctx, &wg)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func TestWithValue(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	ctx := context.WithValue(context.Background(), "number", 9)
	go func() {
		if v := ctx.Value("number"); v != nil {
			n := v.(int)
			fmt.Printf("Square:%d", n*n)
		}
		wg.Done()
	}()

	wg.Wait()
}
