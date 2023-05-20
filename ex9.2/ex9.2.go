package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("에어컨 ON")
	} else if temp <= 3 {
		fmt.Println("히터를 ON")
	} else if temp <= 18 {
		fmt.Println("외출")
	} else {
		fmt.Println("더워")
	}
}
