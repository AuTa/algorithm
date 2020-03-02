/*
 * @lc app=leetcode id=892 lang=golang
 *
 * [892] Surface Area of 3D Shapes
 *
 * https://leetcode.com/problems/surface-area-of-3d-shapes/description/
 *
 * algorithms
 * Easy (57.77%)
 * Likes:    206
 * Dislikes: 266
 * Total Accepted:    15.7K
 * Total Submissions: 27.1K
 * Testcase Example:  '[[2]]'
 *
 * On a N * N grid, we place some 1 * 1 * 1 cubes.
 *
 * Each value v = grid[i][j] represents a tower of v cubes placed on top of
 * grid cell (i, j).
 *
 * Return the total surface area of the resulting shapes.
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[2]]
 * Output: 10
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,2],[3,4]]
 * Output: 34
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [[1,0],[0,2]]
 * Output: 16
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [[1,1,1],[1,0,1],[1,1,1]]
 * Output: 32
 *
 *
 *
 * Example 5:
 *
 *
 * Input: [[2,2,2],[2,1,2],[2,2,2]]
 * Output: 46
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= N <= 50
 * 0 <= grid[i][j] <= 50
 *
 *
 *
 *
 *
 *
 *
 */
package leetcode892

// @lc code=start
func surfaceArea(grid [][]int) int {
	n := len(grid)
	res := 6 * grid[0][0]
	if grid[0][0] > 0 {
		res -= 2 * (grid[0][0] - 1)
	}
	for j := 1; j < n; j++ {
		res += 6*grid[0][j] - 2*min(grid[0][j-1], grid[0][j])
		if grid[0][j] > 0 {
			res -= 2 * (grid[0][j] - 1)
		}
	}
	for i := 1; i < n; i++ {
		res += 6*grid[i][0] - 2*min(grid[i-1][0], grid[i][0])
		if grid[i][0] > 0 {
			res -= 2 * (grid[i][0] - 1)
		}
		for j := 1; j < n; j++ {
			res += 6*grid[i][j] - 2*min(grid[i][j-1], grid[i][j]) - 2*min(grid[i-1][j], grid[i][j])
			if grid[i][j] > 0 {
				res -= 2 * (grid[i][j] - 1)
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
