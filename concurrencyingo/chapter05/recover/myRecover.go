package main

import (
	"log"
	"os"
	"time"
)

func myRecover() {
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

	// 모니터링 대상이 될 고루틴의 시그니처를 파라미터로 받는다.
	// 하트비트 채널을 반환한다.
	type startGoroutineFn func(done <-chan interface{}, pulseInterval time.Duration) (heartbeat <-chan interface{})

	// params: 1. 모니터링 대상 고루틴 타임아웃 시간, 2. 모니터링 할 고루틴을 실행하기 위한 함수
	newSteward := func(timeout time.Duration, ward startGoroutineFn) startGoroutineFn { // startGoroutineFn 반환은 스튜어드 자체도 모니터링 가능하다는 것을 나타낸다. ?
		return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
			heartbeat := make(chan interface{})

			go func() {
				defer close(heartbeat)

				var wardDone chan interface{}
				var wardHeartbeat <-chan interface{}

				startWard := func() { // 모니터링할 고루틴을 시작시킬 클로저 작성
					wardDone = make(chan interface{}) // 중단신호를 보내야 할 경우를 대비해 와드 전용 done 채널을 만들어준다.
					// 이 와드는 아래에서 정의한 와드이다.
					wardHeartbeat = ward(or(wardDone, done), timeout/2) // 스튜어드가 멈추거나, 스튜어드가 와드고루틴을 멈추게 하려고 하는 경우에 두 Done 채널을 사용하여 멈추기 위해 or채널을 사용한다.
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
							close(wardDone) // done채널을 사용해서 ward를 종료하고 새로 시작한다.
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

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	// 와드의 역할
	// 와드는 위에서 정의한 startGoroutineFn 형태와 동일하다. 정확히는 와드의 형태로 startGoroutineFn을 정의한다.
	ward := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("[ward] : Hello, I'm irresponsible! (응답할 수 없습니다)")
		go func() {
			<-done // 아무것도 안하고 취소만 기다린다. (펄스를 보내지 않음)
			log.Println("[ward] : I am halting.(중단합니다)")
		}()
		return nil // 여기서 heartbeat 채널을 반환하고 지속적으로 pulse를 보내줘야 프로그램이 멈추지 않는다. 여기선 멈추게 하기위해 nil을 반환한다.
	}
	// 와드를 담은 스튜어드 function을 획득한다.
	wardWithSteward := newSteward(4*time.Second, ward) // doWork를 시작시키는 스튜어드를 생성. 타임아웃 4초이므로 4초내로 pulse가 없으면 재시작을 시도한다.

	// done채널을 9초 뒤에 닫는다. done 채널이 닫히면 스튜어드와 와드가 모두 종료된다.
	// 흐름상 프로그램을 종료시키기 위해 포함되어있는 코드이다.
	done := make(chan interface{})
	//time.AfterFunc(9*time.Second, func() {
	//	log.Println("[main] : halting steward and ward.(스튜어드와 와드 중단합니다.)")
	//	close(done)
	//})

	for range wardWithSteward(done, 4*time.Second) {
	} // 예제가 멈추는 것을 막기 위해 스튜어드를 시작시키고 펄스들을 순회한다.

	log.Println("Done")
}
