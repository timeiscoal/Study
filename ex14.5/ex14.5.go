package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}

// 사라진 주소를 사용하려고 할 때 댕글링 오류가 발생한다.
// 하지만 go에서는 정상적으로 동작한다.
// 함수가 실행되면 생기는 지역변수, 매개변수들은 스택에 쌓인다.
// Go에서는 탈출분석을 해서, 함수가 생성되는 변수가 밖으로 탈출하는지, 아닌지 확인
// 그래서 탈출하는 녀석은 heap에 보내진다.
// heap은 쓰임이 없어지면 사라진다.
// 탈출 분석을 통해서 스택에,힙에 저장될지 결정된다.
