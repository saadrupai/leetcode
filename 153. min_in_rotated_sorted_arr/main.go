package main

import (
	"fmt"
	"sort"
)

func findMin(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[0]
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(arr))
}
