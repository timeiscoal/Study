package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Print("%")
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println("%")
	}
	fmt.Println()
}
