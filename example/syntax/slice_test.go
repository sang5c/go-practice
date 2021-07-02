package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
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

func TestAppendCopy(t *testing.T) {
	// given
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := append([]int{}, slice1...)

	// when
	slice2 = append(slice2, 6)

	// then
	assert.NotContains(t, slice1, 6)
}

func TestFuncCopy(t *testing.T) {
	// given
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	// when
	slice2 = append(slice2, 6)

	// then
	assert.NotContains(t, slice1, 6)
}

func TestSliceRemoveElement(t *testing.T) {
	// given
	slice := []int{1, 2, 3, 4, 5}
	targetIndex := 3

	// when
	slice = append(slice[:targetIndex], slice[targetIndex+1:]...)

	// then
	assert.NotContains(t, slice, 4)
	assert.Equal(t, 4, len(slice))
}

func TestSliceAppendElement(t *testing.T) {
	// given
	slice := []int{1, 2, 3, 4, 5}
	targetIndex := 3
	targetValue := 20

	// when
	slice = append(slice, 0)
	copy(slice[targetIndex+1:], slice[targetIndex:]) // 한칸씩 뒤로 민다.
	slice[targetIndex] = targetValue

	// then
	assert.Equal(t, []int{1, 2, 3, 20, 4, 5}, slice)
}

func TestSliceSort(t *testing.T) {
	// given
	slice := []int{5, 2, 3, 4, 1}

	// when
	sort.Ints(slice)

	// then
	assert.Equal(t, []int{1, 2, 3, 4, 5}, slice)
}

type User struct {
	Name string
	Age  int
}
type Users []User

func (users Users) Len() int {
	return len(users)
}
func (users Users) Less(i, j int) bool {
	return users[i].Age < users[j].Age // Age를 기준으로 오름차순 정렬
}
func (users Users) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}
func TestStructSliceSort(t *testing.T) {
	// given
	users := Users{
		{"pso", 99},
		{"park", 55},
		{"sang5c", 88},
	}
	expected := Users{
		{"park", 55},
		{"sang5c", 88},
		{"pso", 99},
	}

	// when
	sort.Sort(users)

	// then
	assert.Equal(t, expected, users)
}
