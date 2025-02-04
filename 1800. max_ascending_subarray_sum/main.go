package main

import "fmt"

func maxAscendingSum(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    longest  := nums[0]
    subseqSum := nums[0]
    for i:=1; i<len(nums); i++{
        if nums[i] >nums[i-1]{
            subseqSum+=nums[i]
        }else{
            subseqSum = nums[i]
        }
        if subseqSum > longest{
            longest = subseqSum
        }
    }
    return longest
}

func main(){
	nums:= []int{10,20,30,5,10,50}
	fmt.Println(maxAscendingSum(nums))
}