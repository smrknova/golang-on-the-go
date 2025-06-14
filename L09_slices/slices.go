package main

import "fmt"

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	//uninitialized slice is nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2, 5)
	// capacity -> maximum numbers of elements can fit
	fmt.Println(cap(nums))
	fmt.Println(nums)

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)

	fmt.Println(nums)
	fmt.Println(cap(nums))
}