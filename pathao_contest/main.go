package main

//pay day trip
// func main() {
// 	var numOfGrps int
// 	count := 0
// 	fmt.Scanf("%d", &numOfGrps)
// 	grpMember := make([]int, numOfGrps)
// 	grpMemberMap := make(map[int]int, numOfGrps)
// 	for i := 0; i < numOfGrps; i++ {
// 		fmt.Scanf("%d", &grpMember[i])
// 		if grpMember[i]%4 == 0 {
// 			count++
// 		} else {
// 			grpMemberMap[grpMember[i]]++
// 		}
// 	}

// 	for i := 0; i < numOfGrps; i++ {
// 		if grpMember[i]%4 != 0 {
// 			rem := grpMember[i] % 4
// 			if rem != 2 {
// 				complement := 4 - rem
// 				val1 := grpMemberMap[complement]
// 				val2 := grpMemberMap[rem]
// 				if val1 > 0 && val2 > 0 {
// 					count++
// 					grpMemberMap[complement]--
// 					grpMemberMap[rem]--
// 				} else if rem == 1 {
// 					count += (val2 / 4)
// 					grpMemberMap[rem] = (val2 % 4)
// 				}
// 			} else {
// 				val := grpMemberMap[rem]
// 				if val > 1 {
// 					count += (val / 2)
// 					grpMemberMap[rem] = (val % 2)
// 				}
// 			}
// 		}

// 	}
// 	for key, val := range grpMemberMap {
// 		println(count, key, val)
// 		count += val
// 	}
// 	fmt.Println(count)

// }

// On a pleasant morning, Mr. Sajib and Mr. Samad decided to organize a refreshing trip to Gazipur for the Pathao Pay team. They tasked Mr. Anis with arranging the daylong trip to ensure a memorable and seamless experience for everyone.

// Pathao's employees are divided into several teams, and each team is further divided into smaller groups. To ensure convenience and enjoyment, each group must travel together in the same private car, and a single car can accommodate up to 4 passengers.

// The sizes of these groups, ranging from 1 to 4 members, have already been collected. Now, Mr. Anis needs to determine the minimum number of cars required to transport all the groups while adhering to these rules:

// Each group must stay together in the same car.
// A single car can accommodate multiple groups if their combined size does not exceed 4.
// Help Mr. Anis calculate the minimum number of cars required to transport all the groups.

// Input Format

// The first line of each case will take a single integer g - where g is the number of groups (1 ≤ g ≤ 10⁵) .
// The second line of each case will take g integers p1, p2, ...., pg (1 ≤ pi ≤ 4) represents the size of the i-th group.
// Input terminates at the end of file.
// Constraints

// (1 ≤ g ≤ 10⁵)
// (1 ≤ pi ≤ 4)
// Output Format

// Print a single integer for each test case, representing the minimum number of cars required. Each result must be printed on a separate line.

// Sample Input 0

// 4
// 4 3  2  1
// Sample Output 0

// 3

//bamboo
// Saffiyah loves to play with different DIY projects. Sometimes she uses bamboo sticks for her DIY projects. One day she took some bamboo sticks of the same length and cut them randomly until all parts were at most 50 units long. Now she wants to return the bamboo sticks to their original state, but she forgot how many bamboo sticks she had originally and how long they were originally. Please help her and write a program that calculates the smallest possible original length of those bamboo sticks. All lengths expressed in units are integers greater than zero.

// Input Format

// The input file contains blocks of 2 lines. The first line contains the number of bamboo stick parts after cutting. The second line contains the lengths of those parts separated by the space.

// Constraints

// All integers will be less than 10000 and greater than 0 except the last line of the file contains "0", which denotes the end of test cases.

// Output Format

// The output file contains the smallest possible length of original bamboo sticks, one per line.

// Sample Input 0

// 9
// 5 2 1 5 2 1 5 2 1
// 4
// 1 2 3 4
// 0
// Sample Output 0

// 6
// 5

//discount
// Digital Payments Limited (DPL), a subsidiary of Pathao Limited and widely recognized as Pathao Pay, has emerged as a prominent player in the fintech industry. Signing up for Pathao Pay is quick and hassle-free, where you simply provide your phone number, national ID, and a selfie. While Pathao Pay offers a wide array of features, an exciting new feature is on the horizon: exclusive discounts tailored to their valuable customers based on their names.

// One unique aspect of Pathao Pay is its special emphasis on certain characters. The DPL engineers assign particular importance to the vowels A, E, I, O, and U, collectively known as DPL characters, while all other characters are categorized as non-DPL characters. When reading names from National ID Cards using OCR, sometimes characters are unclear and shown as '?'. These names have uppercase letters (A-Z) and '?' marks. The '?' can be replaced with any letter to complete the name.

// A name (or string) is classified into one of three categories based on the following rules:

// DPL100:
// Contains three consecutive DPL characters, or
// Contains five consecutive non-DPL characters, or
// Meets both conditions.
// DPL00:
// Does not meet any of the above conditions.
// DPL50:
// Can be classified as both DPL100 and DPL00, depending on how '?' is replaced.
// Discount Benefits:

// DPL100 customers: Receive a 100% discount on transactions.
// DPL50 customers: Receive a 50% discount on transactions.
// DPL00 customers: Receive no discount (0%) on transactions.
// The DPL engineers are too lazy to manually classify customer names. If you want to join the DPL family, help them automate the process and identify which customers belong to DPL100, DPL50, or DPL00!

// Input Format

// Input starts with an integer T (≤ 200), denoting the number of test cases.
// Each case begins with a non-empty string with a length no more than 50.
// Constraints

// 1≤∣S∣≤100 (length of the string).
// The string contains only uppercase English letters (A-Z) and '?'.
// Output Format

// For each case of input, you have to print the case number and the result according to the description.

// Sample Input 0

// 4
// SHERLOCK
// SH?RL?CK
// SABA
// MRPCMITTER
// Sample Output 0

// Case 1: 0%
// Case 2: 50%
// Case 3: 0%
// Case 4: 100%

//abir's shoes
// Abir, a passionate settlement officer at Pathao Pay, is responsible for processing a long queue of payment settlements daily. Each settlement arrives with an associated transaction fee, and Abir must prepare a queue of settlements for processing.

// Pathao Pay has specific rules for settlement queues:

// Settlements must be in strictly decreasing order of transaction fees.
// Settlements with a fee of 0 must be ignored.
// Abir faces a challenge: he cannot rearrange the settlements after they arrive. For each settlement, he must make an immediate decision on how to handle it. He has the following three options:

// Add the settlement to the beginning of the queue.
// Add the settlement to the end of the queue.
// Skip the settlement entirely.
// Abir's goal is to create the longest possible sequence of settlements that satisfies the above rules.

// Input Format

// The first line is the number of days for which you have to help Abir. The days follow, one after another; the format of each day is the following:

// Following an integer totalTransaction which is the total transaction number for the day.
// Each of the following totalTransaction line contains a non-negative integer that gives the fees for a transaction. No two transactions have the same fees.
// Constraints

// 0 <= totalTransaction <= 200

// Output Format

// Output a single integer representing the number of transactions in the longest settlement queue that can be formed under the given restrictions. Print each result on a separate line.

// Sample Input 0

// 1
// 7
// 3
// 4
// 2
// 6
// 5
// 0
// 1
// Sample Output 0

// 5

//table tennis
// At Pathao Pay, there's a friendly but intense rivalry between two teams, Team Arif and Team Samad. To decide which team is the best, they’ve organized a series of table tennis matches. Each match is played between one member of Team Arif and one member of Team Samad. However, the records of the matches only show who played against whom, without mentioning which player belongs to which team. Now, the employees at Pathao Pay are curious. They want to figure out the maximum possible number of players in either team based on the match data.

// Input Format

// The first line of input contains an integer T (≤ 10), representing the number of tournaments held.
// Each tournament begins with an integer n (1 <= n <= 10^5), denoting the number of matches scheduled.
// The next n lines each contain two distinct integers u and v (1 ≤ u, v ≤ 20000), indicating that player u will play against player v. No pair of players will be listed more than once.
// Constraints

// T ≤ 10
// 1 <= n <= 10^5
// 1 ≤ u, v ≤ 20000
// Output Format

// For each tournament, print a single line with the tournament number, followed by the maximum possible number of players on either Mr. Arif's team or Mr. Samad's team. Each result should be printed on a separate line.

// Sample Input 0

// 2
// 2
// 1 2
// 2 3
// 3
// 1 2
// 2 3
// 4 2
// Sample Output 0

// Tournament 1: 2
// Tournament 2: 3
