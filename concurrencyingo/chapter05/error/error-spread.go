package main

import (
	"fmt"
	"log"
	"os"
)

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key)) // log 추적이 가능하도록 id를 prefix로 남긴다.
	log.Printf("%#v", err)                           // 로그에는 에러를 남기고
	fmt.Printf("[%v] %v", key, message)              // 사용자에게는 메시지를 반환한다.
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	err := runJob("1")
	if err != nil {
		msg := "there was an unexpected issue; please report this as a bug."
		if _, ok := err.(IntermediateErr); ok {
			msg = err.Error()
		}
		handleError(1, err, msg)
	}
}
