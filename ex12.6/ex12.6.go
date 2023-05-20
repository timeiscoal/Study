package main

import "fmt"

func main() {

	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, "")
		}
		fmt.Println()
	}

}
// 배열 크기 = 타입 크기 * 항목 개수
// [2][5]int = 2 * 5 * 8 = 80

