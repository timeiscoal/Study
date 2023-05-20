package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {

	// 현재 false 가 먼저 연산되면서 &&연산자는 뒷값을
	// 확인하지 않고 그냥 넘어감;;; 뒤에도 false인데..
	// 이런 상황을 조심해야한다.
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}

}

// if 조건문에서 (함수 호출할 경우, 내부 조작이나, 연산
// 을 하는 경우가 조금 위험할 수 있다.
// 그냥 값만 비교 할 수 있게 최대한..!!
// 값을 read만 할 수 있는 펑션, write는 조금..;;
// 무시될 수 있기 떄문에 조금 주의해야 한다.)
// 뭐 값을 검증하는 것만 함수 호출로 사용해라 인듯?
// if 문 소괄호 활용해서 우선 연산과 가독성을 챙긴다.

