package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	game := New()
	assert.NotNil(t, game)
}

func TestIsMatched(t *testing.T) {
	game := New()
	matched, _ := game.Compare(20)
	assert.Equal(t, 0, matched)
}

func TestCountUp(t *testing.T) {
	// given
	game := New()

	// when
	matched, count := game.Compare(12)

	// then
	assert.NotEqual(t, 0, matched)
	assert.Equal(t, 1, count)
}
