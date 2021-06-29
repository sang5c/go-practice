package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unicode/utf8"
	"unsafe"
)

func TestBackQuote(t *testing.T) {
	str1 := `1234\n`
	str2 := "1234\n"
	fmt.Println(str1, str2)

	assert.NotEqual(t, str1, str2)
}

func TestLen(t *testing.T) {
	str1 := "abc" // UTF-8, 3(length) * 1byte
	str2 := "가나다" // UTF-8, 3(length) * 3byte

	assert.Equal(t, 3, len(str1))
	assert.Equal(t, 9, len(str2))
}

func TestStrLen(t *testing.T) {
	str1 := "hello"
	str2 := "안녕!"
	str3 := "hello안녕"

	assert.Equal(t, 5, len([]rune(str1)))
	assert.Equal(t, 3, len([]rune(str2)))
	assert.Equal(t, 7, len([]rune(str3)))
	assert.Equal(t, 7, utf8.RuneCountInString(str3))
}

func TestStringStruct(t *testing.T) {
	str1 := "하이"
	str2 := str1
	assert.Equal(t, "하이", str2)

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	assert.Equal(t, stringHeader1, stringHeader2)

	slice := []byte(str1)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	assert.NotEqual(t, stringHeader1, sliceHeader)
}

func TestConcat(t *testing.T) {
	str1 := "hi"
	temp := str1
	str1 += "hello"

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&temp))

	assert.NotEqual(t, stringHeader1, stringHeader2)
}
