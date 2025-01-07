package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	longest := 1
	for i := 0; i < len(nums); i++ {
		count := 1
		temp := nums[i]
		for j := 0; j < len(nums); j++ {
			if nums[j] == temp+1 {
				count++
				temp++
			}
		}
		if longest < count {
			longest = count
		}
	}
	return longest
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	fmt.Println(longestConsecutive(nums))
}

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
