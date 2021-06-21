package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperator_NewOperator(t *testing.T) {
	operator := NewOperator("+")
	actual := operator.Calc(1, 2)

	assert.Equal(t, 3, actual, "EQUAL")
}
