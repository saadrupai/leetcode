package main

import "fmt"

// can also be implemented in nlogn by sorting the array first, then check consecutive value
func containsDuplicate(arr []int) bool {
	numMap := make(map[int]bool, 0)

	for _, item := range arr {
		if _, ok := numMap[item]; ok {
			return true
		} else {
			numMap[item] = true
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(arr))
}
