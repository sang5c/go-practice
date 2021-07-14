package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	line, err := ReadFile(filename)
	if err == nil {
		t.Fail()
	}

	err = WriteFile(filename, "this is writefile")
	if err != nil {
		fmt.Println("파일 생성 실패", err)
		t.Fail()
	}

	line, err = ReadFile(filename)
	if err != nil {
		fmt.Println("파일 읽기 실패", err)
		t.Fail()
	}

	fmt.Println("파일 내용: ", line)
}

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, fmt.Errorf("제곱근은 양수여야 함. f:%g", f)
		//return 0, fmt.Errorf("제곱근은 양수여야 함. f:%g", f)
		str := fmt.Sprintf("제곱근은 양수여야 함. f:%g", f)
		return 0, errors.New(str)
	}
	return math.Sqrt(f), nil
}

func Test2(t *testing.T) {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("err %v\n", err)
		//t.Fail()
	}
	fmt.Printf("sqrt(-2) = %v\n", sqrt)
}
