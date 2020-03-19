/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (53.12%)
 * Likes:    3110
 * Dislikes: 97
 * Total Accepted:    476K
 * Total Submissions: 887.9K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
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
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 */
package leetcode39

// @lc code=start
import "sort"

func combinationSum(candidates []int, target int) [][]int {
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
		p := make([]int, len(path), len(path)+1)
		copy(p, path)
		res = dfs(remain-num, candidates[i:], append(p, num), res)
	}
	return res
}

// @lc code=end
