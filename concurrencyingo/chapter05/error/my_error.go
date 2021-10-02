package main

import (
	"fmt"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func (err MyError) Error() string {
	return err.Message
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,                            // 가장 낮은 수준의 에러(실제 방생한 에러)
		Message:    fmt.Sprintf(messagef, msgArgs), // 사용자가 읽고 이해할 수 있도록 메시지를 작성하여 담는다.
		StackTrace: string(debug.Stack()),          // 현재(명령이 실행된) 고루틴의 콜스택을 담는다.
		Misc:       make(map[string]interface{}),
	}
}
