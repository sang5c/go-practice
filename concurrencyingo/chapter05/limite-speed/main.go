package main

import (
	multi_limiter "concurrencyingo/chapter05/limite-speed/multi-limiter"
	"concurrencyingo/chapter05/limite-speed/token_bucket"
)

func main() {
	token_bucket.Do()
	token_bucket.Do2()
	multi_limiter.Do()
	multi_limiter.Do2()
}
