/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (45.48%)
 * Likes:    1384
 * Dislikes: 55
 * Total Accepted:    291.9K
 * Total Submissions: 636.2K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 *
 * Each number in candidates may only be used once in the combination.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 */
package leetcode40

import "sort"

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)
	return dfs(target, candidates, path, res)
}

func dfs(remain int, candidates, path []int, res [][]int) [][]int {
	if remain == 0 {
		res = append(res, path)
		return res
	}
	for i, num := range candidates {
		if num > remain {
			return res
		}
		if i > 0 && num == candidates[i-1] { // 如果当前的数和前一个数相同，则不用重复查找了
			continue
		}
		p := make([]int, len(path), len(path)+1)
		copy(p, path)
		res = dfs(remain-num, candidates[i+1:], append(p, num), res)
	}
	return res
}

// @lc code=end
