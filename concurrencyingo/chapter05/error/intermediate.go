package main

import "os/exec"

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

// file path가 실행 가능한지 확인하고 실행한다. 실행이 불가능한 경우 에러를 반환한다.
func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGloballyExec(jobBinPath) // 잘못된 경로를 전달하기 때문에 에러가 발생한다.
	if err != nil {
		return IntermediateErr{
			wrapError(
				err,
				"cannot run job %q: requisite binaries not available",
				id,
			)}
	} else if isExecutable == false { // 프로그램 실행은 가능하나 권한이 없는 경우를 에러로 정의한다.
		return wrapError(
			nil,
			"cannot run job %q: requisite binaries are not executable",
			id,
		)
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}
