package datastructure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]string)
	m["test"] = "value"

	for key, val := range m {
		fmt.Println(key, val)
	}

	assert.Equal(t, "value", m["test"])
}

func TestDelete(t *testing.T) {
	m := make(map[string]string)
	m["test"] = "value"
	m["a"] = "b"

	delete(m, "test")
	_, ok := m["test"]

	assert.False(t, ok)
}

func TestHash(t *testing.T) {
	const M = 10
	hash := func(d int) int {
		return d % M
	}

	m := [M]int{}
	m[hash(23)] = 10
	m[hash(259)] = 50

	fmt.Println(m)
}
