package token_bucket

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"os"
	"sync"
)

func Open2() *APIConnection2 {
	return &APIConnection2{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

type APIConnection2 struct {
	rateLimiter *rate.Limiter
}

func (a *APIConnection2) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// 여기서 작업하는 척 한다.
	return nil
}

func (a *APIConnection2) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	// 여기서 작업하는 척 한다.
	return nil
}

func Do2() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile : %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot resolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}
