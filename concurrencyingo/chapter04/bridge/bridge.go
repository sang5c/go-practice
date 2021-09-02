package main

import "fmt"

var (
	orDone = func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:

					}
				}
			}
		}()
		return valStream
	}
)

func main() {
	bridge := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				stream := make(<-chan interface{})
				select {
				case <-done:
					return
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				}
				for val := range orDone(done, stream) { // TODO: 왜 ordone을 사용한거지?
					select {
					case <-done:
						return
					case valStream <- val:
					}
				}
			}
		}()
		return valStream
	}

	// 여러 채널을 사용하는 환경을 만들기 위해 사용한듯 함.
	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}

}
