package datastructure

import (
	"container/ring"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRing(t *testing.T) {
	r := ring.New(5)
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	assert.Equal(t, int('A'), r.Value)
	assert.Equal(t, int('E'), r.Prev().Value)
	assert.Equal(t, r.Value, r.Move(r.Len()).Value)
}
