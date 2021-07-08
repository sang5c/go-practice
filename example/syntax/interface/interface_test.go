package _interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterface(t *testing.T) {
    // given
	student := Student{"박", 99}
	var stringer Stringer
	stringer = student

	// when
	str := stringer.String()

    // then
	assert.Equal(t, "안녕 99, 박", str)
}

func TestTypeAssertion(t *testing.T) {
	student := &Student{"박", 99}
	var stringer Stringer
	stringer = student

	// when
	_, ok := stringer.(*Student)

	// then
	assert.True(t, ok)

}