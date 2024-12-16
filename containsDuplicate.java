// import java.util.HashMap;
import java.util.HashSet;
import java.util.Scanner;

public class containsDuplicate{
    public static boolean solution(int[] arr){

        // method:1
        // HashMap<Integer,Boolean> solutionMap = new HashMap<>();
        // for(int i=0; i<arr.length; i++){
        //     if (solutionMap.containsKey(arr[i])){
        //         return true;
        //     }else{
        //         solutionMap.put(arr[i], true);
        //     }
        // }
        // return false;

        // method:2
        HashSet<Integer> numSet = new HashSet<>();

        for (int n : arr){
            if (numSet.contains(n)){
                return true;
            }else{
                numSet.add(n);
            }
        }
        return false;
    } 

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int size = scanner.nextInt();

        int[] num = new int[size];
        for (int i=0; i<num.length; i++){
            num[i] = scanner.nextInt();
        }
        scanner.close();
        boolean res = solution(num);

        System.out.println(res);
    }
}


// 217. Contains Duplicate
// Solved
// Easy
// Topics
// Companies
// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

 

// Example 1:

// Input: nums = [1,2,3,1]

// Output: true

// Explanation:

// The element 1 occurs at the indices 0 and 3.

// Example 2:

// Input: nums = [1,2,3,4]

// Output: false

// Explanation:

// All elements are distinct.

// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]

// Output: true

 

// Constraints:

// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109