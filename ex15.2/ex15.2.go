package main

import "fmt"

func main() {

	poet1 := "안녕하세요 반갑습니다.\n두번째줄\n세번째줄\n"

	poet2 := `안녕하세요 반갑습니다
두번째줄
세번째줄`

	fmt.Println(poet1)
	fmt.Println(poet2)
}
