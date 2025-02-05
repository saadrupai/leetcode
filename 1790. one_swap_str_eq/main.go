package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {

    if s1 == s2{
        return true
    }
    if len(s1) != len(s2){
        return false
    }
    mismatchedIdx := []int{}
    for i:=0; i<len(s1); i++{
        if s1[i] != s2[i]{
            mismatchedIdx = append(mismatchedIdx, i)
        }

        if len(mismatchedIdx)>2{
            return false
        }
    }

    if len(mismatchedIdx)==2{
        i,j:= mismatchedIdx[0],mismatchedIdx[1]
        return s1[i]==s2[j] && s1[j]==s2[i] 
    }

   return false

}


func main(){
	s1 := "bank"
	s2 := "kanb"
	fmt.Println(areAlmostEqual(s1,s2))
}


// 1790. Check if One String Swap Can Make Strings Equal
// Easy
// Topics
// Companies
// Hint
// You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.

// Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.

 

// Example 1:

// Input: s1 = "bank", s2 = "kanb"
// Output: true
// Explanation: For example, swap the first character with the last character of s2 to make "bank".
// Example 2:

// Input: s1 = "attack", s2 = "defend"
// Output: false
// Explanation: It is impossible to make them equal with one string swap.
// Example 3:

// Input: s1 = "kelb", s2 = "kelb"
// Output: true
// Explanation: The two strings are already equal, so no string swap operation is required.
 

// Constraints:

// 1 <= s1.length, s2.length <= 100
// s1.length == s2.length
// s1 and s2 consist of only lowercase English letters.