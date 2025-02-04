package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	alphaNum := []rune{}
	lowerS := strings.ToLower(s)
	for _, char := range lowerS {
		if (char >= 'a' && char <= 'z') ||
			(char >= '0' && char <= '9') {
			alphaNum = append(alphaNum, char)
		}
	}
	str := string(alphaNum)
	println(str)
	for i, j := 0, len(alphaNum)-1; i < j; i, j = i+1, j-1 {
		alphaNum[i], alphaNum[j] = alphaNum[j], alphaNum[i]
	}
	rstr := string(alphaNum)
	println(rstr)
	if str == rstr {
		return true
	} else {
		return false
	}
}

func main() {
	s := "A man, a plan, a canal: Panama"

	fmt.Println(isPalindrome(s))
}

// Code
// Code
// Code Sample
// Testcase
// Testcase
// Test Result
// 125. Valid Palindrome
// Solved
// Easy
// Topics
// Companies
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// Constraints:

// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters.
