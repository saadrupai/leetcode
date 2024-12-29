package main

import "fmt"

func twoSum(arr []int, target int) []int {
	numMap := make(map[int]int, 0)

	for i, item := range arr {
		complement := target - item
		if j, ok := numMap[complement]; ok {
			return []int{i, j}
		} else {
			numMap[item] = i
		}
	}
	return []int{}
}

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}
