package go_goroutine

import (
	"fmt"
	"time"
)

func TestHandler() {

	fmt.Println("testing")
	TestGo()
}

// 코어 개수가 3개 이상이 되지 않으면, 이 세 고루틴을 동시에 실행시킬 코어가 부족해서, 동시에 실행되지 않는다.
// 마치 동시에 실행되는 것처럼 보인다.

func TestGo() {

	// 새로운 고루틴 생성
	go PrintHangul()
	go PrintNumber()

	// 함수 실행에 필요한 시간을 알 수 있다면 좋겠지만, 아래와 같이 함수 실행에 필요한 시간을 알지 못한다면
	// 함수는 실행 중 종료가 될 것이다.

	// 1초 대기 (3초 대기는 해야 모든 go루틴 실행이 가능하다)
	time.Sleep(1 * time.Second)

}

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}

}

func PrintNumber() {

	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}

}
