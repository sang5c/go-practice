package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSlice(t *testing.T) {
	var slice1 []int
	slice2 := []int{}
	slice3 := make([]int, 0)

	assert.Nil(t, slice1)    // 빈 슬라이스는 Nil이다.
	assert.NotNil(t, slice2) // Nil이 아닌 empty slice이다.
	assert.NotNil(t, slice3)
}

func TestCapLen(t *testing.T) {
	slice := make([]int, 3, 5)

	assert.Equal(t, 5, cap(slice))
	assert.Equal(t, 3, len(slice))
}

func TestAppend(t *testing.T) {
	var slice []int
	slice = append(slice, 100, 99)

	assert.Contains(t, slice, 100)
}

func TestSlicing(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[0:2]
	slice3 := slice2[1:]

	//slice1 = append(slice1, 6)
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3[0] = 100
	slice3 = append(slice3, 200)
	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))
}
