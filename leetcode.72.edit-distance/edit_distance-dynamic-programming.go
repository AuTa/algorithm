/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 *
 * https://leetcode.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (41.17%)
 * Likes:    3296
 * Dislikes: 50
 * Total Accepted:    240.9K
 * Total Submissions: 570.7K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * Given two words word1 and word2, find the minimum number of operations
 * required to convert word1 to word2.
 *
 * You have the following 3 operations permitted on a word:
 *
 *
 * Insert a character
 * Delete a character
 * Replace a character
 *
 *
 * Example 1:
 *
 *
 * Input: word1 = "horse", word2 = "ros"
 * Output: 3
 * Explanation:
 * horse -> rorse (replace 'h' with 'r')
 * rorse -> rose (remove 'r')
 * rose -> ros (remove 'e')
 *
 *
 * Example 2:
 *
 *
 * Input: word1 = "intention", word2 = "execution"
 * Output: 5
 * Explanation:
 * intention -> inention (remove 't')
 * inention -> enention (replace 'i' with 'e')
 * enention -> exention (replace 'n' with 'x')
 * exention -> exection (replace 'n' with 'c')
 * exection -> execution (insert 'u')
 *
 *
 */
package leetcode72

// @lc code=start
func minDistance(word1 string, word2 string) int {
	l2 := len(word2)
	var d []int = make([]int, l2+1) // 用单行代替多行
	for j := 0; j <= l2; j++ {
		d[j] = j
	}
	for i := 1; i <= len(word1); i++ {
		old := i - 1 // 左上角的值
		d[0] = i     // 左边的值
		for j := 1; j <= len(word2); j++ {
			temp := d[j]
			if word1[i-1] == word2[j-1] { // 字符串索引从 0 开始
				d[j] = old
			} else {
				d[j] = minOfThree(d[j], d[j-1], old) + 1
			}
			old = temp
		}
	}
	return d[len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minOfThree(x, y, z int) int {
	return min(min(x, y), z)
}

// @lc code=end

// 分治 Time Limit Exceeded
func minDistance01(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	if l2 == 0 {
		return l1 // b 为空，等于删除 a 的所有字符。
	}
	if l1 == 0 {
		return l2 // a 为空，等于插入 b 的所有字符。
	}
	if word1[l1-1] == word2[l2-1] {
		return minDistance01(word1[:l1-1], word2[:l2-1]) // a[i] == b[j] 时，d[i][j] = d[i-1][j-1]
	}
	return minOfThree(
		minDistance01(word1[:l1-1], word2)+1,        // 插入 a[i]，d[i][j] = d[i-1][j]+1
		minDistance01(word1, word2[:l2-1])+1,        // 删除 b[j]，d[i][j] = d[i][j-1]+1
		minDistance01(word1[:l1-1], word2[:l2-1])+1, // a[i] 替换成 b[j]，d[i][j] = d[i-1][j-1]+1
	)
}

// minDistanceDP 用动态规划以空间换时间。
// d[i][j] 总是和 d[i-1][j], d[i][j-1], d[i-1][j-1] 相关。
func minDistanceDP(word1 string, word2 string) int {
	var d [][]int // d[i][j] 表示 a[i]->b[j] 需要的距离
	for i := 0; i <= len(word1); i++ {
		d = append(d, make([]int, len(word2)+1))
		d[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		d[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] { // 字符串索引从 0 开始
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = minOfThree(d[i-1][j], d[i][j-1], d[i-1][j-1]) + 1
			}
		}
	}
	return d[len(word1)][len(word2)]
}
