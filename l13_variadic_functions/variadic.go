package main

import "fmt"

// use interface{} for any type in place of int
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
	// fmt.Println(1, 2, 3, 5, 86, "hello")
	// result := sum(3, 4, 5, 6)

	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)
}