package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p1 == p3 : %v\n", p1 == p3)
	fmt.Println(unsafe.Sizeof(p1))
}
