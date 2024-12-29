package main

import (
	"fmt"
)

// sorting
// func isAnagram(s, t string) bool {

// 	if len(s) != len(t) {
// 		return false
// 	}
// 	bytesS := []byte(s)
// 	byteT := []byte(t)
// 	sort.Slice(bytesS, func(i, j int) bool {
// 		return bytesS[i] < bytesS[j]
// 	})
// 	sort.Slice(byteT, func(i, j int) bool {
// 		return byteT[i] < byteT[j]
// 	})

// 	sortedS := string(bytesS)
// 	sortedT := string(byteT)

// 	return sortedS == sortedT
// }

// hash table
// func isAnagram(s, t string) bool {

// 	if len(s) != len(t) {
// 		return false
// 	}

// 	charMap := make(map[rune]uint)

// 	for _, ch := range s {
// 		charMap[ch]++
// 	}

// 	for _, ch := range t {
// 		if count, ok := charMap[ch]; ok {
// 			if count > 0 {
// 				charMap[ch]--
// 			} else {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}

// 	return true
// }

func isAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	char := make([]int, 26)

	for _, ch := range s {
		v := int(ch - 'a')
		char[v]++
	}

	for _, ch := range t {
		v := int(ch - 'a')
		char[v]--
	}

	for _, ch := range char {
		if ch != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "aacc"
	t := "ccac"
	fmt.Println(isAnagram(s, t))
}

// Given an array of strings strs, group the
// anagrams
//  together. You can return the answer in any order.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]

// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Explanation:

// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
// Example 2:

// Input: strs = [""]

// Output: [[""]]

// Example 3:

// Input: strs = ["a"]

// Output: [["a"]]
