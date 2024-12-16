import java.util.Arrays;
import java.util.HashMap;
import java.util.Scanner;

class twoSum {
    public static int[] solution(int[] nums, int target) {

        // bruteforce
        // for (int i=0; i<nums.length; i++){
        //     for (int j=i+1; j<nums.length; j++){
        //         if (nums[i]+nums[j] == target){
        //             return new int[]{i,j};
        //         }
        //     }
        // }

        // using hashmap
        HashMap<Integer,Integer> numMap = new HashMap<>();

        for (int i=0; i<nums.length ; i++){
            int complement = target - nums[i];
            if (numMap.containsKey(complement)){
                return new int[]{numMap.get(complement), i};
            }
            numMap.put(nums[i], i);

        }

        return null;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int size = scanner.nextInt();

        int[] nums = new int[size];

        int target = scanner.nextInt();

        for (int i=0; i<size; i++){
            nums[i] = scanner.nextInt();
        }

        scanner.close();

        int[] res = solution(nums, target);
        if (res != null) {
            System.out.println("Indices: " + Arrays.toString(res));
        } else {
            System.out.println("No two sum solution found.");
        }
    }
}


// 1. Two Sum
// Solved
// Easy
// Topics
// Companies
// Hint
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

 

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]
 

// Constraints:

// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109