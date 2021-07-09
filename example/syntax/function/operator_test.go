package function

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuncVar(t *testing.T) {
	operator := getOperator("+")
	result := operator(2, 3)
	assert.Equal(t, 5, result)
}
