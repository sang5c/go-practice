package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
		Inner:      err, // 가장 낮은 수준의 에러(실제 방생한 에러)
		Message:    fmt.Sprintf(messagef, msgArgs),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

// lowlevel 모듈
type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{wrapError(err, err.Error())}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}

// intermediate
type IntermediateErr struct {
	error
}

// before
//func runJob(id string) error {
//	const jobBinPath = "/bad/job/binary"
//	isExecutable, err := isGloballyExec(jobBinPath)
//	if err != nil {
//		return err
//	} else if isExecutable == false {
//		return wrapError(nil, "job binary is not executable")
//	}
//	return exec.Command(jobBinPath, "--id="+id).Run()
//}

func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return IntermediateErr{
			wrapError(
				err,
				"cannot run job %q: requisite binaries not available",
				id,
			)}
	} else if isExecutable == false {
		return wrapError(
			nil,
			"cannot run job %q: requisite binaries are not executable",
			id,
		)
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
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
