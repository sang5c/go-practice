package main

import "os"

type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path) // FileInfo 객체를 얻는다. 파일 이름, 정보 등이 담겨있다.
	if err != nil {
		return false, LowLevelErr{
			wrapError(err, err.Error()),
		}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}
