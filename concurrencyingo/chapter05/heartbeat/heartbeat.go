package main

import (
	"fmt"
	"time"
)

var (

	// pulseInterval은 하트비트 간격을 의미한다.
	doWork = func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		// 결과 채널과 함께 하트비트 채널을 반환한다.
		// 사용자는 결과채널에 값이 들어오지 않아도 무한정 기다리지 않고 하트비트가 도착하는 것을 보고 정상 동작중임을 알 수 있다.
		heartbeat := make(chan interface{})
		results := make(chan time.Time)

		go func() {
			defer close(heartbeat)
			defer close(results)

			pulse := time.Tick(pulseInterval)       // Tick 함수를 사용하여 pulse 채널에 지속적으로 신호를 보낸다.
			workGen := time.Tick(2 * pulseInterval) // 완료되었을때 발생하는 티커. 하트비트를 보기위해 하트비트보다 큰 값인 두배를 사용했다.

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default: // heartbeat 채널을 listening 하지 않고 있을 때를 대비하여 존재하는 default이다. 값을 안읽어가면 쓰기가 막혀서 진행되지 않는다.
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse: // 작업 완료 전 pulse가 발생하면 다시 입력해주는 과정이 필요하다.
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse: // 여기도 마찬가지. 펄스를 계속 발생시킨다.
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()

		return heartbeat, results
	}

	doWorkPanic = func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{})
		results := make(chan time.Time)
		go func() {
			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-pulse:
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}
)

func main() {
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { // 10초뒤 done채널을 닫는다.
		close(done)
	})

	const timeout = 2 * time.Second // 타임아웃 처리 시간.
	heartbeat, results := doWork(done, timeout/2)
	//heartbeat, results := doWorkPanic(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		case <-time.After(timeout): // 일정 시간동안 값을 받지 못한다면 타임아웃 처리
			fmt.Println("worker goroutine is not healthy!")
			return
		}
	}
}
