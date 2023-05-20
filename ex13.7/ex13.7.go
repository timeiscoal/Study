package main

import (
	"fmt"
	"unsafe"
)

// type User struct {
// 	A int8 // 1
// 	B int  // 8
// 	C int8 // 1
// 	D int  // 8
// 	E int8 // 1
// 	// 19byte => 40byte
// }

type User struct {
	A int8 // 1
	C int8 // 1
	E int8 // 1
	B int  // 8
	D int  // 8

}

func main() {

	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
	// 40바이트
}

// 0 8 16 24 32 40byte
// 0 8 16 24byte
