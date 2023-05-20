package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{22.1, 23.2, 24.1, 25.1, 26.1}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

}

// x:=[...]은 배열이고, 아직 크기가 정해지지 않았다.
// x:= [...]int{10,20,30} 이랑 []int{10,20,30} 서로 다르다.
