/*
 * @lc app=leetcode id=1329 lang=golang
 *
 * [1329] Sort the Matrix Diagonally
 *
 * https://leetcode.com/problems/sort-the-matrix-diagonally/description/
 *
 * algorithms
 * Medium (78.01%)
 * Likes:    95
 * Dislikes: 34
 * Total Accepted:    5.8K
 * Total Submissions: 7.4K
 * Testcase Example:  '[[3,3,1,1],[2,2,1,2],[1,1,1,2]]'
 *
 * Given a m * n matrix mat of integers, sort it diagonally in ascending order
 * from the top-left to the bottom-right then return the sorted array.
 *
 *
 * Example 1:
 *
 *
 * Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
 * Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
 *
 *
 *
 * Constraints:
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 1 <= m, n <= 100
 * 1 <= mat[i][j] <= 100
 *
 */
package leetcode1329

// @lc code=start
func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	for i := 0; i < m-1; i++ {
		for j := 0; j < m-1-i; j++ { // 将最大的值交换到最下面行，然后除去最下面行后继续排序
			for k := 0; k < n-1; k++ {
				if mat[j][k] > mat[j+1][k+1] {
					mat[j][k], mat[j+1][k+1] = mat[j+1][k+1], mat[j][k]
				}
			}
		}
	}
	return mat
}

// @lc code=end
