package main

import "fmt"

func countBits(n int) []int {
	resp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		div := i
		binNum := 0
		if i == 0 {
			binNum = 0
		}
		for div != 0 {
			carry := div % 2
			binNum = binNum*10 + carry
			div = div / 2
		}
		println(i, binNum)
		resp[i] = binNum
	}
	return resp
}

func main() {
	ip := 5
	fmt.Println(countBits(ip))
}

// Code
// Testcase
// Test Result
// Test Result
// 338. Counting Bits
// Easy
// Topics
// Companies
// Hint
// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

// Example 1:

// Input: n = 2
// Output: [0,1,1]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// Example 2:

// Input: n = 5
// Output: [0,1,1,2,1,2]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101

// Constraints:

// 0 <= n <= 105

// Follow up:

// It is very easy to come up with a solution with a runtime of O(n log n). Can you do it in linear time O(n) and possibly in a single pass?
// Can you do it without using any built-in function (i.e., like __builtin_popcount in C++)?
