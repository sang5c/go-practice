package multi_limiter

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sync"
	"time"
)

func Open2() *APIConnection2 {
	return &APIConnection2{
		apiLimit: MultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 2),
			rate.NewLimiter(Per(10, time.Minute), 10),
		),
		diskLimit: MultiLimiter(
			rate.NewLimiter(rate.Limit(1), 1),
		),
		networkLimit: MultiLimiter(
			rate.NewLimiter(Per(3, time.Second), 3),
		),
	}
}

type APIConnection2 struct {
	networkLimit,
	diskLimit,
	apiLimit RateLimiter
}

func (a *APIConnection2) ReadFile(ctx context.Context) error {
	if err := MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx); err != nil {
		return err
	}
	// 여기서 작업하는 척 한다.
	return nil
}

func (a *APIConnection2) ResolveAddress(ctx context.Context) error {
	if err := MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx); err != nil {
		return err
	}
	// 여기서 작업하는 척 한다.
	return nil
}

func Do2() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open2()
	var wg sync.WaitGroup

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("canno ReadFile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}
