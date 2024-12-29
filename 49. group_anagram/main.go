package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(arr []string) [][]string {

	anagramMap := make(map[string][]string, 0)
	result := make([][]string, 0)

	for _, str := range arr {
		bytesStr := []byte(str)
		sort.Slice(bytesStr, func(i, j int) bool {
			return bytesStr[i] < bytesStr[j]
		})
		sortedStr := string(bytesStr)
		if _, ok := anagramMap[sortedStr]; ok {
			anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
		} else {
			anagramMap[sortedStr] = []string{str}
		}
	}

	for _, anagram := range anagramMap {
		result = append(result, anagram)
	}

	return result
}

func main() {
	arr := []string{"a"}

	fmt.Println(groupAnagrams(arr))
}
