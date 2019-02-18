/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (24.86%)
 * Total Accepted:    272.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 *
 *
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 * Note:
 *
 *
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 *
 * Example 2:
 *
 *
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the precedeng element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 *
 *
 * Example 3:
 *
 *
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 *
 *
 * Example 4:
 *
 *
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore
 * it matches "aab".
 *
 *
 * Example 5:
 *
 *
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 *
 *
 */
package leetcode0010

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	if p[len(p)-1] == '*' {
		// 是否匹配零个
		ans := isMatch(s, p[:len(p)-2])
		if ans {
			return true
		}
		// s 最后一个字符能匹配，去掉后重新匹配
		if s != "" && (p[len(p)-2] == '.' || s[len(s)-1] == p[len(p)-2]) {
			ans = isMatch(s[:len(s)-1], p)
		}
		return ans
	}
	if s == "" {
		return false
	}
	if p[len(p)-1] != '.' && s[len(s)-1] != p[len(p)-1] {
		return false
	}
	return isMatch(s[:len(s)-1], p[:len(p)-1])
}
