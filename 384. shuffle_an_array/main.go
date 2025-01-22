package main

import "math/rand"

// optimized

// type Solution struct {
// 	original *[]int
// 	shuffle  *[]int
// 	r        *rand.Rand
// }

// func Constructor(nums []int) Solution {
// 	original := make([]int, len(nums))
// 	shuffle := make([]int, len(nums))
// 	copy(original, nums)
// 	copy(shuffle, nums)
// 	sol := Solution{original: &original, shuffle: &shuffle, r: rand.New(rand.NewSource(time.Now().UnixNano()))}
// 	return sol
// }

// func (this *Solution) Reset() []int {
// 	return *this.original
// }

// func (this *Solution) Shuffle() []int {
// 	this.r.Shuffle(len(*this.shuffle), func(i, j int) {
// 		shuf := *this.shuffle
// 		shuf[i], shuf[j] = shuf[j], shuf[i]
// 	})
// 	return *this.shuffle
// }

// my soln

type Solution struct {
	Nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		Nums: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.Nums
}

func (this *Solution) Shuffle() []int {
	numArr := make([]int, len(this.Nums))
	copy(numArr, this.Nums)
	for i, _ := range numArr {
		temp := rand.Intn(i + 1)
		numArr[i], numArr[temp] = numArr[temp], numArr[i]
	}

	return numArr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

// 384. Shuffle an Array
// Solved
// Medium
// Topics
// Companies
// Hint
// Given an integer array nums, design an algorithm to randomly shuffle the array. All permutations of the array should be equally likely as a result of the shuffling.

// Implement the Solution class:

// Solution(int[] nums) Initializes the object with the integer array nums.
// int[] reset() Resets the array to its original configuration and returns it.
// int[] shuffle() Returns a random shuffling of the array.

// Example 1:

// Input
// ["Solution", "shuffle", "reset", "shuffle"]
// [[[1, 2, 3]], [], [], []]
// Output
// [null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]

// Explanation
// Solution solution = new Solution([1, 2, 3]);
// solution.shuffle();    // Shuffle the array [1,2,3] and return its result.
//                        // Any permutation of [1,2,3] must be equally likely to be returned.
//                        // Example: return [3, 1, 2]
// solution.reset();      // Resets the array back to its original configuration [1,2,3]. Return [1, 2, 3]
// solution.shuffle();    // Returns the random shuffling of array [1,2,3]. Example: return [1, 3, 2]

// Constraints:

// 1 <= nums.length <= 50
// -106 <= nums[i] <= 106
// All the elements of nums are unique.
// At most 104 calls in total will be made to reset and shuffle.
