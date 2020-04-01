/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 *
 * https://leetcode.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (29.67%)
 * Likes:    2020
 * Dislikes: 117
 * Total Accepted:    231.8K
 * Total Submissions: 776.1K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 *
 * Each element in the array represents your maximum jump length at that
 * position.
 *
 * Your goal is to reach the last index in the minimum number of jumps.
 *
 * Example:
 *
 *
 * Input: [2,3,1,1,4]
 * Output: 2
 * Explanation: The minimum number of jumps to reach the last index is 2.
 * ⁠   Jump 1 step from index 0 to 1, then 3 steps to the last index.
 *
 * Note:
 *
 * You can assume that you can always reach the last index.
 *
 */
package leetcode45

// @lc code=start
func jump(nums []int) int {
	steps := 0
	period := [2]int{0, 1}
	for period[1] < len(nums) {
		maxPos := 0
		for i := period[0]; i < period[1]; i++ {
			if i+nums[i] > maxPos {
				maxPos = i + nums[i]
			}
		}
		period[0], period[1] = period[1], maxPos+1
		steps++
	}
	return steps
}

// @lc code=end
