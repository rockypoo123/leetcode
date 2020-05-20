//link: https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

/*
In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.

Return the element repeated N times.

 

Example 1:

Input: [1,2,3,3]
Output: 3
Example 2:

Input: [2,1,2,5,3,2]
Output: 2
Example 3:

Input: [5,1,5,2,5,3,5,4]
Output: 5
 

Note:
4 <= A.length <= 10000
0 <= A[i] < 10000
A.length is even
*/

/*
Result:
Runtime: 32 ms, faster than 48.21% of Go online submissions for N-Repeated Element in Size 2N Array.
Memory Usage: 6.3 MB, less than 28.57% of Go online submissions for N-Repeated Element in Size 2N Array.
*/

func repeatedNTimes(A []int) int {
    m := make(map[int]int)
    
    for i := 0; i < len(A); i ++ {
        if m[A[i]] > 0 {
            return A[i]
        }
        m[A[i]] ++
    }
    
    return -1
}