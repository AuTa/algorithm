/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 *
 * https://leetcode.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (49.48%)
 * Likes:    492
 * Dislikes: 66
 * Total Accepted:    48.9K
 * Total Submissions: 98.8K
 * Testcase Example:  '13'
 *
 * Given an integer n, return 1 - n in lexicographical order.
 *
 * For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].
 *
 * Please optimize your algorithm to use less time and space. The input size
 * may be as large as 5,000,000.
 *
 */
package leetcode386

// @lc code=start
func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	for i := 1; i <= 9; i++ {
		res = lexical(res, i, n)
	}
	return res
}

// 递归向后补充数字
func lexical(res []int, current, n int) []int {
	if current > n {
		return res
	}
	res = append(res, current)
	for i := 0; i <= 9; i++ {
		res = lexical(res, current*10+i, n)
	}
	return res
}

// @lc code=end
