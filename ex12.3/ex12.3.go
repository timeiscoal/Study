package main

import "fmt"

func main() {

	nums := [...]int{10, 20, 30, 40, 50} //[5]int{10,20..}
	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	var t [5]float64 = [5]float64{1.1, 1.2, 1.3, 1.4, 1.5}

	t[3] = 1.44

	for i, v := range t {
		fmt.Println(i, v)
	}
}
