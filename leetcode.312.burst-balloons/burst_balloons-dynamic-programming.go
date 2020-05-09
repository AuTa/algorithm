/*
 * @lc app=leetcode id=312 lang=golang
 *
 * [312] Burst Balloons
 *
 * https://leetcode.com/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (50.66%)
 * Likes:    2116
 * Dislikes: 58
 * Total Accepted:    90.1K
 * Total Submissions: 177.9K
 * Testcase Example:  '[3,1,5,8]'
 *
 * Given n balloons, indexed from 0 to n-1. Each balloon is painted with a
 * number on it represented by array nums. You are asked to burst all the
 * balloons. If the you burst balloon i you will get nums[left] * nums[i] *
 * nums[right] coins. Here left and right are adjacent indices of i. After the
 * burst, the left and right then becomes adjacent.
 *
 * Find the maximum coins you can collect by bursting the balloons wisely.
 *
 * Note:
 *
 *
 * You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can
 * not burst them.
 * 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
 *
 *
 * Example:
 *
 *
 * Input: [3,1,5,8]
 * Output: 167
 * Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  -->
 * []
 * coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
 *
 */
package leetcode312

// @lc code=start
func maxCoins(nums []int) int {
	n := len(nums)

	p := make([]int, 0, n+2)
	p = append(append(append(p, 1), nums...), 1)

	dp := make([][]int, n+2)
	dp[n+1] = make([]int, n+2)
	for i := len(nums); i >= 0; i-- {
		dp[i] = make([]int, n+2)
		for j := i + 1; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+p[i]*p[k]*p[j])
			}
		}
	}
	return dp[0][n+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
