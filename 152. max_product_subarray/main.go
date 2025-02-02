package main

import "fmt"

func maxProduct(nums []int) int {
	max := nums[0]
	min := nums[0]
	resp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max, min = min, max
		}

		maxProd := nums[i] * max
		if maxProd > nums[i] {
			max = maxProd
		} else {
			max = nums[i]
		}

		minProd := nums[i] * min
		if minProd < nums[i] {
			min = minProd
		} else {
			min = nums[i]
		}

		if max > resp {
			resp = max
		}

	}

	return resp
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}
