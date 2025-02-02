package main

import "fmt"

func maxProduct(nums []int) int {
	max := nums[0]
	min := nums[0]
	resp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max, min = min, max
		}

		maxProd := nums[i] * max
		if maxProd > nums[i] {
			max = maxProd
		} else {
			max = nums[i]
		}

		minProd := nums[i] * min
		if minProd < nums[i] {
			min = minProd
		} else {
			min = nums[i]
		}

		if max > resp {
			resp = max
		}

	}

	return resp
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}

// 152. Maximum Product Subarray
// Solved
// Medium
// Topics
// Companies
// Given an integer array nums, find a
// subarray
//  that has the largest product, and return the product.

// The test cases are generated so that the answer will fit in a 32-bit integer.

// Example 1:

// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
// Example 2:

// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

// Constraints:

// 1 <= nums.length <= 2 * 104
// -10 <= nums[i] <= 10
// The product of any subarray of nums is guaranteed to fit in a 32-bit integer.
