/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (25.68%)
 * Likes:    5696
 * Dislikes: 690
 * Total Accepted:    787.3K
 * Total Submissions: 3.1M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 */
package leetcode15

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 { // 当前数字大于 0，那么后续的和肯定都大于 0
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1         // 左边的数从当前数后一个开始
		r := len(nums) - 1 // 右边的数从最后一个数开始
		for l < r {        // 遍历
			if nums[r] < 0 { // 右边数比 0 小的话，和肯定一直小于 0
				break
			}
			if s := nums[i] + nums[l] + nums[r]; s > 0 { // 和大于 0，则右边向前移一位
				r -= 1
			} else if s < 0 { // 和小于 0，则左边向后移一位
				l += 1
			} else { // 同时移位
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				r -= 1
				l += 1
			}
		}
	}
	return res
}

// @lc code=end
