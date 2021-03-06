//link: https://leetcode.com/problems/number-of-equivalent-domino-pairs/

/*
Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if and only if either (a==c and b==d), or (a==d and b==c) - that is, one domino can be rotated to be equal to another domino.

Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to dominoes[j].

 

Example 1:

Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
Output: 1
 

Constraints:

1 <= dominoes.length <= 40000
1 <= dominoes[i][j] <= 9
*/

/*
Result:
Runtime: 3708 ms, faster than 6.00% of Go online submissions for Number of Equivalent Domino Pairs.
Memory Usage: 7.7 MB, less than 100.00% of Go online submissions for Number of Equivalent Domino Pairs.
*/

func numEquivDominoPairs(dominoes [][]int) int {
    ans := 0
    
    for i := 0; i < len(dominoes); i ++ {
        for j := i + 1; j < len(dominoes); j ++ {
            if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) || (dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0]) {
                ans ++
            }
        }
    }
    return ans
}