package subtests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTeardownParallel(t *testing.T) {
	// <setup code>
	// This Run will not return until its parallel subtests complete.
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", func(t *testing.T) {
			t.Error("FAIL")
		})
		t.Run("Test2", func(t *testing.T) {
			assert.True(t, false)
		})
		t.Run("Test3", func(t *testing.T) {
			assert.True(t, false)
		})
	})
	// <tear-down code>
}
