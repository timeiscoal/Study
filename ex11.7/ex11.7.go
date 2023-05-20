package main

import "fmt"

// 플래그 변수를 활용한다.
// 플래그 변수를 통해서 언제 빠져나갈지 정하는 것.
// 그런데 중첩 반복문이 많거나 문장이 복잡하면 플레그 좀 어렵
// 레이블도 있다 중첩 반복문 빠져나가기
func main() {

	a := 1
	b := 2
	flag := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	fmt.Println(a, b, a*b)
}
