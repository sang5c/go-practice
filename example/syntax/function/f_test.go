package function

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestVarargs(t *testing.T) {
	sum := func(nums ...int) int {
		result := 0
		for _, num := range nums {
			result += num
		}
		return result
	}
	assert.Equal(t, 5, sum(1, 2, 2))
}

func TestLiteral(t *testing.T) {
	result := func() int {
		return 10
	}()

	assert.Equal(t, 10, result)
}

type fn func() int

func TestCapture(t *testing.T) {
	f := make([]fn, 3)
	result := make([]int, 3)

	t.Run("변수를 참조로 가져오면", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			f[i] = func() int {
				return i
			}
		}
		for i, function := range f {
			result[i] = function()
		}

		t.Run("for문이 종료했을때의 값이 된다.", func(t *testing.T) {
			assert.Equal(t, []int{3, 3, 3}, result)
		})
	})
	t.Run("변수를 복사해서 가져오면", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			v := i
			f[i] = func() int {
				return v
			}
		}
		for i, function := range f {
			result[i] = function()
		}

		t.Run("각 호출시점의 for문 인덱스가 된다.", func(t *testing.T) {
			assert.Equal(t, []int{0, 1, 2}, result)
		})
	})
}

func TestFileHandle(t *testing.T) {
	text := "hello world"

	type Writer func(string)
	writeHello := func(writer Writer) {
		writer(text)
	}

	file, _ := os.Create("test.txt")
	defer file.Close()
	defer os.Remove("test.txt")

	// dependency injection
	writeHello(func(msg string) {
		fmt.Fprint(file, msg)
	})

	readFile, _ := os.ReadFile("test.txt")

	assert.Equal(t, text, string(readFile))
}
