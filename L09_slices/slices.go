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

	// var nums = make([]int, 2, 5)
	// capacity -> maximum numbers of elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)

	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// var nums2 = make([]int, 0, 5)
	// nums2 = append(nums2, 1)
	// nums2 = append(nums2, 2)
	// fmt.Println(nums2)


	nums := []int{}

	nums = append(nums, 1)
	nums = append(nums, 2)

	nums[0] = 3
	nums[1] = 5
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// // copy function
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator

	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// slices
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2))

	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)

}