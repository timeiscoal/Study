package main

import "fmt"

func main() {

	//정수 변수 a 선언 및 초기화
	var a int = 500
	//포인터 변수 p 선언 (메모리 주소를 값으로 가지는 타입)
	var p *int
	fmt.Println(p) // nil 이네 주소값 없음

	// 포인터 변수 p에 a의 메모리 주소를 넣어준다.
	p = &a

	fmt.Printf("p의 값 : %p\n", p)
	fmt.Println(&a)
	fmt.Println(p, *p)
	fmt.Printf("p가 가리키는 메모리 값 : %d\n", *p)

	*p = 100
	fmt.Printf("a의 값 : %d\n", a)
}
