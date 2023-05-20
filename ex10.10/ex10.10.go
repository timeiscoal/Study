package main

import "fmt"

func main() {

	a := 3
	switch a {
	case 1:
		fmt.Println("1")
		// break
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough // 밑으로 떨어진다.
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("a != 1,2,3,4")
	}
}

// result = 3 4
