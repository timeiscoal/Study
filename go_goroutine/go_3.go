package go_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

// 경쟁상황
// 예측은 2000이 되어야 하지만, 실제 출력되는 값의 범위는 1800~1990 사이
// 두 고루틴이 같은 같은 변수의 값을 사용하면서 발생하는 문제.

func TestGo3() {
	testMutex()
	testMutex2()

}

func testMutex() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU를 사용한다.

	var data = []int{} // 슬라이스 생성

	// 고루틴 익명함수
	// 변수 data 에 1을 1000번 추가하기.
	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보한다.
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1)
			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))
}

func testMutex2() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU를 사용한다.

	var data = []int{}

	// mutex := &sync.Mutex{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()           // 뮤텍스를 잠궈서 , data 슬라이스를 보호
			data = append(data, 1) // 추가 이후
			mutex.Unlock()         // 잠금해제
			runtime.Gosched()      // 다른 고루틴이 CPU를 사용할 수 있도록 양보한다.
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()
			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))

}
