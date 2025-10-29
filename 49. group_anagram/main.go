package main

import (
	"fmt"
	"sort"
)

// func groupAnagrams(strs []string) [][]string {
//     groups := make(map[[26]int][]string)

//     for _, word := range strs {
//         var freq [26]int
//         for _, ch := range word {
//             freq[ch-'a']++
//         }
//         groups[freq] = append(groups[freq], word)
//     }

//     res := make([][]string, 0, len(groups))
//     for _, group := range groups {
//         res = append(res, group)
//     }
//     return res
// }

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
