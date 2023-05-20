package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

// *Data Data type의 주소값을 받았다.
// 주소
func ChangeData(arg *Data) {
	//func ChangeData(arg Data) {

	(*arg).value = 999
	// arg.value = 999
	arg.data[100] = 999
}

func main() {

	var data Data
	//ChangeData(data)
	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}

// result 0 0
// 값이 바뀌지 않았다. 해결 방법 포인터로 가자.

// changeData 함수를 호출 할때 어떤 알규먼트를 넣게 된다.
// 이때 알규먼트는(우변)값으로 쓰인다. r-value
// 이 말의 의미는 ChanageData(arg Data)에서 arg = data와 같다.
// 즉 복사되서 서로 다른 변수로 메모리 공간에 저장된다.

// var a int
// var b int
// a = b , 두 공간은 서로 다르다. 들어간 값은 같아도!
// 결론은 두 변수의 공간은 서로 다르다! copy.

// 메모리 주소값 8바이트만 복사할 수 있어서 그 만큼 개꿀!
// 성능이 그 만큼 좋아진다.
// 원래 1608바이트 전부 다 복사해야함;;;

// 구조체 포인터 추기화

// 1번 방법
// var data Data
// var p *Data = &data
// data라는 메모리 공간이 있고,
// 변수 p에다가 메모리 공간 주소 값을 저장한다.

// 2번 방법 , 구조체 생성자
// var p *Data = &Data{} // 데이터는 없지만 공간은 있다.이름만 없을 뿐 &Data{} 데이터 구조체
// 어떤 데이터 타입의 공간이 있다.
// 이 데이터 타입의 공간 주소를 p가 가지고 있다.

// 인스턴스
// 인스턴스는 메모리 공간에 할당된 데이터의 실체
// var p *Data = &Data{} Data 인스턴스 하나가 만들어졌고,
// 포인터 변수 p가 가리킨다.

// 3개의 인스턴스 값만 같을 뿐!

// var data1 Data
// var data2 Data = data1
// var data3 Data = data1

// new() 내장함수 서로 같다, 초기화하는 방법
// p1 := &Data{} // 각 필드 값들을 넣어서 초기화 가능
// var p2 = new(Data) // 각 필드 값 넣어서 초기화 불가능

//인스턴스는 언제 사라지나?
// 아무도 찾지 않을 때 사라진다.

type User struct {
	Age int
}

func TestFunc() {
	u := &User{} //u 포인터 변수 선언, 인스턴스 생성
	u.Age = 30   // 값 대입
	fmt.Println(u)
} // 소속된 중괄호가 닫힐 때 이제 사라진다.
// 힙에 저장되었을 경우에는 GC로 사라진다.
