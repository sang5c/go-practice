package method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithdrawPointer(t *testing.T) {
	a := &account{100, "Joe", "Park"}
	t.Run("30을 더하면", func(t *testing.T) {
		a.withdrawPointer(30)
		t.Run("잔고는 70이다", func(t *testing.T) {
			assert.Equal(t, 70, a.balance)
		})
	})
}

func TestWithdrawValueNotChangeBalance(t *testing.T) {
    // given
	a := &account{100, "Joe", "Park"}

    // when
	a.withdrawValue(20)

    // then
    assert.Equal(t, 100, a.balance)
}

func TestWithdrawValueReturn(t *testing.T) {
	// given
	a := &account{100, "Joe", "Park"}

	// when
	b := a.withdrawReturnValue(20)

	// then
	assert.Equal(t, 80, b.balance)
}