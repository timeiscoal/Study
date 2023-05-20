package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	// 구조체 초기화 : 모든 필드 값들은 기본 값으로 초기화
	// var house House
	var house House = House{
		"관악구",
		100,
		99,
		"아파트",
	}
	fmt.Println(house)

	house.Address = "서울특별시"
	house.Size = 10
	house.Price = 30
	house.Category = "주택"

	fmt.Printf("%v\n", house)

}
