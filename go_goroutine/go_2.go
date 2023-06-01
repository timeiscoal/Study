package go_goroutine

import (
	"fmt"
	"sync"
)

// waitGroup 객체

var wg sync.WaitGroup

func SumAtoB(a, b int) {

	sum := 0
	for i := a; i <= b; i++ {

		sum += i
	}

	fmt.Printf("%d 부터 %d 의 합계 => %d\n", a, b, sum)

	wg.Done() // 작업이 완료 됨
}

// 작업의 횟수 = 고루틴 개수 => 실행가능
// 작업의 횟수 > 고루틴 개수 => 실행불가능 (데드락 발생)
// 작업의 횟수 < 고루틴 개수 => 실행가능 : 단 작업의 개수 만큼만 동작함
// ex) wg.Add(10) , go SumAtoB -> 실행결과는 하나

func TestGo2() {

	wg.Add(10) // 총 작업의 횟수를 설정한다

	for i := 0; i < 10; i++ {
		go SumAtoB(1, 10000)
	}

	defer fmt.Println("작업완료")

	wg.Wait() // 모든 작업이 완료되기를 기다린다

}
