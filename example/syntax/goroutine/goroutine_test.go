package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 0; i < +5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func TestGo(t *testing.T) {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
	defer fmt.Println()
}
