package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	game := New()
	assert.NotNil(t, game)
}
