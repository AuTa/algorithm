/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 *
 * https://leetcode.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (29.58%)
 * Likes:    826
 * Dislikes: 168
 * Total Accepted:    185.1K
 * Total Submissions: 625.8K
 * Testcase Example:  '1'
 *
 * Given a positive integer, return its corresponding column title as appear in
 * an Excel sheet.
 *
 * For example:
 *
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "A"
 *
 *
 * Example 2:
 *
 *
 * Input: 28
 * Output: "AB"
 *
 *
 * Example 3:
 *
 *
 * Input: 701
 * Output: "ZY"
 *
 */
package main

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	if x := n % 26; x == 0 {
		return convertToTitle(n/26-1) + "Z"
	} else {
		s := byte(x) + 'A' - 1
		ss := convertToTitle(n / 26)
		return string(append([]byte(ss), s))
	}
}
