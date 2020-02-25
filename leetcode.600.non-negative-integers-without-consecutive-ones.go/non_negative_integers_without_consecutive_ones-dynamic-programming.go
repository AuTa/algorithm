/*
 * @lc app=leetcode id=600 lang=golang
 *
 * [600] Non-negative Integers without Consecutive Ones
 *
 * https://leetcode.com/problems/non-negative-integers-without-consecutive-ones/description/
 *
 * algorithms
 * Hard (33.50%)
 * Likes:    309
 * Dislikes: 66
 * Total Accepted:    10.7K
 * Total Submissions: 31.9K
 * Testcase Example:  '1'
 *
 * Given a positive integer n, find the number of non-negative integers less
 * than or equal to n, whose binary representations do NOT contain consecutive
 * ones.
 *
 * Example 1:
 *
 * Input: 5
 * Output: 5
 * Explanation:
 * Here are the non-negative integers
 *
 *
 * Note:
 * 1
 *
 */
package leetcode600

// @lc code=start
func findIntegers(num int) int {
	return find(0, 0, 0, num)
}

func find(k uint, current, append, num int) int {
	if current > num {
		return 0
	}
	s := 1 << k
	if s > num { // 指某个值一直补 0，直到不能补为止，才统计为 1 次
		return 1
	}
	if append == 0 { // 如果补的是 0，则前面可以再补一个 0 或 1
		return find(k+1, current, 0, num) + find(k+1, current+s, 1, num)
	}
	if append == 1 { // 如果补的是 1，则前面只能补一个 0
		return find(k+1, current, 0, num)
	}
	return 0
}

// @lc code=end
