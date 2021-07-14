package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func divide(a, b int) {
	if b == 0 {
		panic("0으로 나눌 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() { divide(1, 0) })
	assert.PanicsWithValue(t, "0으로 나눌 수 없습니다", func() { divide(1, 0) })
	assert.PanicsWithError(t, "error!", func() { panic(errors.New("error!")) })
}

func TestPanicWithoutAssert(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// The following is the code under test
	divide(1, 0)
}

func panicRecover() {
	defer func() {
		recover()
	}()
	panic("PANIC!")
}

func TestRecover(t *testing.T) {
	assert.NotPanics(t, panicRecover)
}
