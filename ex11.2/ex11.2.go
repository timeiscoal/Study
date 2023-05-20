package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	for {
		//time.Sleep 은 쓰레드를 멈추는 함수 1초
		//무한 루프
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}
