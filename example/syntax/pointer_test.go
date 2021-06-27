package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointerAddress(t *testing.T) {
	var value1 = 10
	var value2 = 10

	var pValue1 *int = &value1
	var pValue2 *int = &value2

	if pValue1 == pValue2 {
		fmt.Printf("%p %p", pValue1, pValue2)

	}
	assert.Equal(t, pValue1, pValue2)
	assert.Same(t, pValue1, pValue2) // Pointer는 Same으로 비교해야 한다.
}

func TestPointerAddress2(t *testing.T) {
	a := 10
	b := 20

	pA := &a
	pB := &b
	pC := &a

	assert.Equal(t, pA, pC)
	assert.NotEqual(t, pA, pB)

	fmt.Println(pA, pB, pC)
}
