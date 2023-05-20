package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4
	Score float64 // 8
	// 12
}

func main() {

	var a int = 45
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a))
}
