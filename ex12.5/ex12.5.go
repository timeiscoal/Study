package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println("---------------------------")

	for i, v := range b {
		fmt.Printf("a[%d] = %d", i, v)
	}

	fmt.Println("---------------------------")

	// 복사
	b = a
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}

// 복사가 하나씩 옮기는 것이 아니라, 복사한다.
// 항상 같은 공간의 크기가 같아야 한다.
// 공간 그 자체 전체 모두를 복사하기 떄문에 공간 크기가 반드시 같아야한다.
// 크다고 담을 수 없음.
