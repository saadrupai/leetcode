import java.util.Arrays;
import java.util.Scanner;

public class validAnagram{
    public static boolean solution(String s, String t){
        if (s.length() != t.length()){
            return false;
        }
        char[] charArr = s.toCharArray();
        Arrays.sort(charArr);
        s = new String(charArr);

        charArr = t.toCharArray();
        Arrays.sort(charArr);
        t = new String(charArr);
        
        if (s.equals(t)) {
            return true;
        }

        return false;
    } 

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        String firstString = scanner.nextLine();
        String secondString = scanner.nextLine();

        scanner.close();

        boolean res = solution(firstString, secondString);

        System.out.println(res);
    }
}


// 242. Valid Anagram
// Solved
// Easy
// Topics
// Companies
// Given two strings s and t, return true if t is an 
// anagram
//  of s, and false otherwise.

 

// Example 1:

// Input: s = "anagram", t = "nagaram"

// Output: true

// Example 2:

// Input: s = "rat", t = "car"

// Output: false

 

// Constraints:

// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.
