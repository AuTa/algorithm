/*
 * @lc app=leetcode id=1262 lang=golang
 *
 * [1262] Greatest Sum Divisible by Three
 *
 * https://leetcode.com/problems/greatest-sum-divisible-by-three/description/
 *
 * algorithms
 * Medium (43.02%)
 * Likes:    246
 * Dislikes: 7
 * Total Accepted:    8.3K
 * Total Submissions: 19.2K
 * Testcase Example:  '[3,6,5,1,8]'
 *
 * Given an array nums of integers, we need to find the maximum possible sum of
 * elements of the array such that it is divisible by three.
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [3,6,5,1,8]
 * Output: 18
 * Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum
 * divisible by 3).
 *
 * Example 2:
 *
 *
 * Input: nums = [4]
 * Output: 0
 * Explanation: Since 4 is not divisible by 3, do not pick any number.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,3,4,4]
 * Output: 12
 * Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum
 * divisible by 3).
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 4 * 10^4
 * 1 <= nums[i] <= 10^4
 *
 *
 */
package leetcode1262

// @lc code=start
func maxSumDivThree(nums []int) int {
	dp := [3]int{0, -1000, -1000}
	for _, n := range nums {
		temp := dp
		for i, a := range temp {
			m := (i + n) % 3
			dp[m] = max(temp[m], a+n)
		}
	}
	return dp[0]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
