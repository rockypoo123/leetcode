//link: https://leetcode.com/problems/count-largest-group/

/*
Given an integer n. Each number from 1 to n is grouped according to the sum of its digits. 

Return how many groups have the largest size.

 

Example 1:

Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]. There are 4 groups with largest size.
Example 2:

Input: n = 2
Output: 2
Explanation: There are 2 groups [1], [2] of size 1.
Example 3:

Input: n = 15
Output: 6
Example 4:

Input: n = 24
Output: 5
 

Constraints:
1 <= n <= 10^4
*/

/*
Result:
Runtime: 4 ms, faster than 68.97% of Go online submissions for Count Largest Group.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Count Largest Group.
*/

func countLargestGroup(n int) int {
    m := make(map[int]int)
    max := 0
    for i := 1; i <= n; i ++ {
        count := 0
        j := i
        for j != 0 {
            count += (j % 10)
            j /= 10
        }
        m[count] ++
        if m[count] > max {
            max = m[count]
        }
    }
    ans := 0
    for _, v := range m {
        if v == max {
            ans ++
        }
    }
    return ans
}