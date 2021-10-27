package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func recoverWithClosure() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil // 재귀 탈출 1
		case 1:
			return channels[0] // 재귀 탈출 2
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...): // 재귀
				}
			}
		}()
		return orDone
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

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

	type startGoroutineFn func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (heartbeat <-chan interface{}) // 모니터링, 재시작 대상이 될 고루틴의 시그니처

	newSteward := func(
		timeout time.Duration, // 모니터링 대상 고루틴 타임아웃 시간
		startGoroutine startGoroutineFn, // 모니터링 할 고루틴을 실행하기 위한 함수
	) startGoroutineFn { // 스튜어드 자체도 모니터링 가능하다는 것을 나타낸다.
		return func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} {
			heartbeat := make(chan interface{})

			go func() {
				defer close(heartbeat)

				var wardDone chan interface{}
				var wardHeartbeat <-chan interface{}

				startWard := func() { // 모니터링할 고루틴을 시작시킬 클로저 작성
					wardDone = make(chan interface{})                             // 중단신호를 보내야 할 경우를 대비해 와드 전용 done 채널을 만들어준다.
					wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2) // 스튜어드가 멈추거나, 스튜어드가 와드고루틴을 멈추게 하려고 하는 경우에 두 Done 채널을 사용하여 멈추기 위해 or채널을 사용한다.
				}
				startWard()

				pulse := time.Tick(pulseInterval)

			monitorLoop:
				for {
					timeoutSignal := time.After(timeout)

					for { // 스튜어드가 자체적으로 펄스를 보낼 수 있도록 하는 내부 루프
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case <-wardHeartbeat: // 하트비트가 오면 계쏙 모니터링
							continue monitorLoop
						case <-timeoutSignal: // 타임아웃 설정한 시간 내로 펄스를 못받으면 와드를 멈추고 새로 시작한다. 로그를 남긴 후 와드와 모니터링은 계속됨
							log.Println("[steward] : ward unhealthy; restarting (와드가 비정상, 재시작)")
							close(wardDone)
							startWard()
							continue monitorLoop
						case <-done:
							return
						}
					}
				}
			}()

			return heartbeat
		}
	}

	// -=================================

	doWorkFn := func(
		done <-chan interface{},
		intList ...int,
	) (startGoroutineFn, <-chan interface{}) {
		intChanStream := make(chan (<-chan interface{}))
		intStream := bridge(done, intChanStream)
		doWork := func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} {
			intStream := make(chan interface{})
			heartbeat := make(chan interface{})
			go func() {
				defer close(intStream)
				select {
				case intChanStream <- intStream:
				case <-done:
					return
				}

				pulse := time.Tick(pulseInterval)

				for {
				valueLoop:
					for _, intVal := range intList {
						if intVal < 0 {
							log.Printf("negative value: %v\n", intVal)
							return
						}

						for {
							select {
							case <-pulse:
								select {
								case heartbeat <- struct{}{}:
								default:
								}

							case intStream <- intVal:
								continue valueLoop
							case <-done:
								return
							}
						}
					}
				}
			}()
			return heartbeat
		}
		return doWork, intStream

	}

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	done := make(chan interface{})
	defer close(done)

	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 5)
	doWorkWithSteward := newSteward(1*time.Millisecond, doWork)
	doWorkWithSteward(done, 1*time.Hour)

	for intVal := range take(done, intStream, 6) {
		fmt.Printf("Received: %v\n", intVal)
	}

	log.Println("Done")
}
