/*
 * @lc app=leetcode id=672 lang=golang
 *
 * [672] Bulb Switcher II
 *
 * https://leetcode.com/problems/bulb-switcher-ii/description/
 *
 * algorithms
 * Medium (50.52%)
 * Likes:    111
 * Dislikes: 786
 * Total Accepted:    11.3K
 * Total Submissions: 22.3K
 * Testcase Example:  '1\n1'
 *
 * There is a room with n lights which are turned on initially and 4 buttons on
 * the wall. After performing exactly m unknown operations towards buttons, you
 * need to return how many different kinds of status of the n lights could be.
 *
 * Suppose n lights are labeled as number [1, 2, 3 ..., n], function of these 4
 * buttons are given below:
 *
 *
 * Flip all the lights.
 * Flip lights with even numbers.
 * Flip lights with odd numbers.
 * Flip lights with (3k + 1) numbers, k = 0, 1, 2, ...
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: n = 1, m = 1.
 * Output: 2
 * Explanation: Status can be: [on], [off]
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: n = 2, m = 1.
 * Output: 3
 * Explanation: Status can be: [on, off], [off, on], [off, off]
 *
 *
 *
 *
 * Example 3:
 *
 *
 * Input: n = 3, m = 1.
 * Output: 4
 * Explanation: Status can be: [off, on, off], [on, off, on], [off, off, off],
 * [off, on, on].
 *
 *
 *
 *
 * Note: n and m both fit in range [0, 1000].
 *
 */
package leetcode672

// @lc code=start
func flipLights(n int, m int) int {
	if m == 0 {
		return 1
	}
	switch n {
	case 1:
		return 2
	case 2:
		if m == 1 {
			return 3
		}
		return 4
	default:
		switch m {
		case 1:
			return 4
		case 2:
			return 7
		default:
			return 8
		}
	}
}

// @lc code=end
