package main

import (
	"fmt"
)

func maxArea(height []int) int {
	area := 0
	i, j := 0, len(height)-1
	for i < j {
		width := j - i

		var heig int
		if height[i] > height[j] {
			heig = height[j]
			j--
		} else {
			heig = height[i]
			i++
		}
		currentArea := int(width) * heig
		// println(width, heig, currentArea)
		if currentArea > area {
			area = currentArea
		}
	}
	return area
}

func main() {
	height := []int{2, 1}
	fmt.Println(len(height))
	fmt.Println(maxArea(height))
}
