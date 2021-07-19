package goroutine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestWg(t *testing.T) {

	var wg sync.WaitGroup
	var results = make([]bool, 10)

	SumAtoB := func(a, b, index int) {
		defer wg.Done()
		results[index] = true

		sum := 0
		for i := 0; i <= b; i++ {
			sum += i
		}
		fmt.Printf("%d부터 %d까지 합계는 %d\n", a, b, sum)
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000, i)
	}
	wg.Wait()

	assert.NotContains(t, results, false)
}

type Account struct {
	Balance int
}

func DepositAndWithDraw(account *Account, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("balance should not be negative value: %d", account.Balance))
	}

	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func TestRunConcurrent(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	account := &Account{0}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithDraw(account, &mutex)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func TestDivideArea(t *testing.T) {

	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
