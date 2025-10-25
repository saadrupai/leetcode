package main

import "fmt"

type Stack struct {
	data []rune
}

func (s *Stack) push(val rune) {
	s.data = append(s.data, val)
}

func (s *Stack) pop() rune {
	if len(s.data) == 0 {
		return 0
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *Stack) isEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func isValid(s string) bool {
	stack := Stack{}
	parenPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack.push(ch)
		}

		if ch == ')' || ch == '}' || ch == ']' {
			top := stack.pop()
			if top != parenPairs[ch] {
				return false
			}
		}
	}

	return stack.isEmpty()

}

func main() {
	s := "({)}"

	if isValid(s) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}

// 20. Valid Parentheses
// Solved
// Easy
// Topics
// premium lock icon
// Companies
// Hint
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"

// Output: true

// Example 2:

// Input: s = "()[]{}"

// Output: true

// Example 3:

// Input: s = "(]"

// Output: false

// Example 4:

// Input: s = "([])"

// Output: true

// Example 5:

// Input: s = "([)]"

// Output: false

// Constraints:

// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.
