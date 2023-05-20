package main

import "fmt"

type User struct {
	Age int
	Id  string
}

type VIPUser struct {
	User     // 내장된 필드
	VIPLevel int
	price    int
}

func main() {

	user := User{23, "김강민"}
	vip := VIPUser{
		User{30, "황철순"},
		3,
		250,
	}

	fmt.Println(user)
	fmt.Println(vip)
	fmt.Println(vip.price, vip.VIPLevel, vip.Id, vip.User.Age)

}
