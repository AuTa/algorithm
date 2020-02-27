/*
 * @lc app=leetcode id=1005 lang=golang
 *
 * [1005] Maximize Sum Of Array After K Negations
 *
 * https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/
 *
 * algorithms
 * Easy (50.74%)
 * Likes:    234
 * Dislikes: 29
 * Total Accepted:    20.8K
 * Total Submissions: 41K
 * Testcase Example:  '[4,2,3]\n1'
 *
 * Given an array A of integers, we must modify the array in the following way:
 * we choose an i and replace A[i] with -A[i], and we repeat this process K
 * times in total.  (We may choose the same index i multiple times.)
 *
 * Return the largest possible sum of the array after modifying it in this
 * way.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: A = [4,2,3], K = 1
 * Output: 5
 * Explanation: Choose indices (1,) and A becomes [4,-2,3].
 *
 *
 *
 * Example 2:
 *
 *
 * Input: A = [3,-1,0,2], K = 3
 * Output: 6
 * Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
 *
 *
 *
 * Example 3:
 *
 *
 * Input: A = [2,-3,-1,5,-4], K = 2
 * Output: 13
 * Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 10000
 * 1 <= K <= 10000
 * -100 <= A[i] <= 100
 *
 *
 */
package leetcode1005

// @lc code=start
import (
	"sort"
)

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	for i := 0; i < K; i++ {
		if A[i] < 0 {
			A[i] = -A[i]
		} else {
			if i == 0 {
				A[i] = -A[i]
			} else if A[i] != 0 && (K-i)%2 != 0 {
				if A[i] < A[i-1] {
					A[i] = -A[i]
				} else {
					A[i-1] = -A[i-1]
				}
			}
			break
		}
	}
	r := 0
	for _, a := range A {
		r += a
	}
	return r
}

// @lc code=end
