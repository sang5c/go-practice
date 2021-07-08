package _interface

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕 %d, %s", s.Age, s.Name)
}
