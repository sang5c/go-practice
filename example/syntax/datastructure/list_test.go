package datastructure

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	getElement := func(list *list.List, index int) interface{} {
		if list.Len() < index {
			panic("out of range")
		}
		element := v.Front()
		for i := 0; i < index; i++ {
			element = element.Next()
		}
		return element.Value
	}

	element := getElement(v, 2)
	assert.Equal(t, 3, element)
}

type Queue struct {
	list *list.List
}

func (q *Queue) Push(value interface{}) {
	q.list.PushBack(value)
}

func (q *Queue) Pop() interface{} {
	element := q.list.Front()
	if element == nil {
		return nil
	}
	return q.list.Remove(element)
}

func New() *Queue {
	return &Queue{list.New()}
}

func TestQueue(t *testing.T) {
	queue := New()
	queue.Push(4)
	queue.Push(1)
	queue.Push(2)

	assert.Equal(t, 4, queue.Pop())
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, 2, queue.Pop())
}

type Stack struct {
	list *list.List
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() interface{} {
	element := s.list.Back()
	if element == nil {
		return nil
	}
	return s.list.Remove(element)
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 3, stack.Pop())
}
