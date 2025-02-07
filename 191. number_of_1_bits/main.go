package main

import "fmt"

func hammingWeight(n int) int {
	if n == 0 {
		return 0
	}
	div := n / 2
	carry := n % 2
	count := 1
	for div != 0 {
		if carry != 0 {
			count++
		}
		carry = div % 2

		div = div / 2
	}
	return count
}

func main() {
	a := 2147483645
	fmt.Println(hammingWeight(a))
}
