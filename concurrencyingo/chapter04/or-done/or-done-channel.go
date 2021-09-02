package main

import "fmt"

func main() {
	done := make(chan interface{})
	myChan := make(chan interface{})

	// 적용 전
	for val := range myChan {
		// do
		fmt.Println(val)
	}

	// 적용 전, 확장
	// done 채널이 닫혔을때와 myChan이 닫혔을때의 동작을 구분하고 싶다는 요구사항.
loop:
	for {
		select {
		case <-done:
			break loop // 레이블 사용, 복잡도 증가.
		case maybeVal, ok := <-myChan:
			if ok == false {
				return // 혹은 break로 for문을 벗어남
			}
			// val로 무언가 작업
			fmt.Println(maybeVal)
		}
	}

	// 적용 후
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
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

	for val := range orDone(done, myChan) {
		// do something
		fmt.Println(val)
	}
}
