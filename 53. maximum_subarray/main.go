package main

import (
	"fmt"
)

// optimized
func maxSubArray(nums []int) int {
	max := nums[0]
	currentSum := 0

	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		if currentSum > max {
			max = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}
	return max
}

//bruteforce
// func maxSubArray(nums []int) int {
// 	max := nums[0]

// 	for i := 0; i < len(nums); i++ {
// 		currentSum := 0
// 		for j := i; j < len(nums); j++ {
// 			currentSum += nums[j]
// 			if currentSum > max {
// 				max = currentSum
// 			}
// 		}
// 	}
// 	return max
// }

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

// 53. Maximum Subarray
// Medium
// Topics
// Companies
// Given an integer array nums, find the
// subarray
//  with the largest sum, and return its sum.

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.
// Example 2:

// Input: nums = [1]
// Output: 1
// Explanation: The subarray [1] has the largest sum 1.
// Example 3:

// Input: nums = [5,4,-1,7,8]
// Output: 23
// Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
