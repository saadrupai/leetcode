package main

import (
	"fmt"
	"sort"
)

func topKFrequent(arr []int, k int) []int {
	freqMap := map[int]int{}

	if len(arr) == 1 {
		return []int{arr[0]}
	}

	type keyValue struct {
		Key   int
		Value int
	}

	keyVal := []keyValue{}
	resp := []int{}

	for _, val := range arr {
		freqMap[val]++
	}

	for key, val := range freqMap {
		keyVal = append(keyVal, keyValue{
			Key:   key,
			Value: val,
		})
	}

	sort.Slice(keyVal, func(i int, j int) bool {
		return keyVal[i].Value > keyVal[j].Value
	})

	for i := range k {
		resp = append(resp, keyVal[i].Key)
	}

	return resp
}

func main() {
	arr := []int{1}
	k := 2
	fmt.Println(topKFrequent(arr, k))
}

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.
