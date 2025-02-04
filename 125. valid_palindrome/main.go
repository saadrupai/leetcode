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
