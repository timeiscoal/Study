package main

import "fmt"

func richMan() bool {
	return true
}

func countFriend() int {
	return 3

}

func main() {
	price := 50000

	if price > 50000 {
		if richMan() {
			fmt.Println("신발 끈 묶기")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 && price <= 50000 {
		if countFriend() > 3 {
			fmt.Println("신발 끈을 묶는다")
		} else {
			fmt.Println("나눠 내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다잉")

	}

}
