/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.71%)
 * Likes:    1664
 * Dislikes: 120
 * Total Accepted:    426.6K
 * Total Submissions: 933K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 *
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */
package leetcode16

import "sort"

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[len(nums)-1]
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if abs(res-target) > abs(s-target) {
				res = s
			}
			if s < target {
				l++
			} else if s > target {
				r--
			} else {
				return s
			}
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end
