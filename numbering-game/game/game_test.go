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
	assert.True(t, game.isMatched(20))
	assert.False(t, game.isMatched(21))
}
