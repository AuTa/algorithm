/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 *
 * https://leetcode.com/problems/nth-digit/description/
 *
 * algorithms
 * Medium (30.99%)
 * Likes:    310
 * Dislikes: 951
 * Total Accepted:    53.8K
 * Total Submissions: 173.6K
 * Testcase Example:  '3'
 *
 * Find the n^th digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8,
 * 9, 10, 11, ...
 *
 * Note:
 * n is positive and will fit within the range of a 32-bit signed integer (n <
 * 2^31).
 *
 *
 * Example 1:
 *
 * Input:
 * 3
 *
 * Output:
 * 3
 *
 *
 *
 * Example 2:
 *
 * Input:
 * 11
 *
 * Output:
 * 0
 *
 * Explanation:
 * The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a
 * 0, which is part of the number 10.
 *
 *
 */
package leetcode400

import "math"

// @lc code=start
func findNthDigit(n int) int {
	base := 9   // 每一轮都是 9 个数字
	digits := 1 // 位数
	number := 1 // 对应的数字
	for n-base*digits > 0 {
		n -= base * digits
		base *= 10
		digits += 1
		number *= 10
	}
	idx := n % digits // 对应数字的第 idx 位是要返回的结果

	if idx == 0 {
		number += n/digits - 1
	} else {
		number += n / digits
		// 去除尾数
		number /= int(math.Pow10(digits - idx))
	}
	return number % 10
}

// @lc code=end
