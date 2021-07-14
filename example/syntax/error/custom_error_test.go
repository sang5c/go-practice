package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CustomError struct {
}

func (error CustomError) Error() string {
	return "커스텀 에러"
}

func TestCustomError(t *testing.T) {
	customError := CustomError{}
	assert.Error(t, customError)
}

// ---

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧음"
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

func TestCustomError2(t *testing.T) {
	password := "myPw"
	err := RegisterAccount("myID", password)
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
			assert.Equal(t, len(password), errInfo.Len)
		}
	} else {
		fmt.Println("회원가입성공")
	}
}
