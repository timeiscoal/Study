package main

import "fmt"

// func star1() {
// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < i+1; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func star2() {
// 	for i := 0; i < 5; i++ {
// 		for j := 4; j > i-1; j-- {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}

// }

func star3() {
	for i := 0; i < 5; i++ {
		for j := 0; j < (2*5 + 1); j++ {
			if j < (5-i) || j > (5+i) {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}

		fmt.Println()
	}
}

func strr4() {
	var a int = 5
	for i := 0; i < a; i++ {
		if i <= a/2 {
			for j := 0; j < (2*a - 1); j++ {
				if j < (a-i) || j > (a+i) {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			}
			fmt.Println()
		} else {
			for j := 0; j < (2*a - 1); j++ {
				if j < i+1 || j > a-i+4 {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			}
			fmt.Println()
		}

	}
}

func main() {
	star3()
	fmt.Println("///////////////")
	strr4()
	// star2()
	// fmt.Println("///////////////")
	// star1()
}
