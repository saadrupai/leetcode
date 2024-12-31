package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	prefix := 1
	resp := make([]int, 4)
	resp[0] = 1
	for i, num := range nums {
		if i+1 == len(nums) {
			break
		}
		prefix = prefix * num
		resp[i+1] = prefix

	}
	postfix := 1
	for i := len(nums) - 1; i > 0; i-- {
		postfix = nums[i] * resp[i]
		resp[i-1] = 
	}
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(arr))
}

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:

// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
