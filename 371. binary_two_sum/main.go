package main

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carryWithShift := (a & b) << 1
		a = sum
		b = carryWithShift
	}
	return a
}

func main() {
	a, b := 1, 2
	fmt.Println(getSum(a, b))
}
