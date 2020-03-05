/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (32.46%)
 * Likes:    1544
 * Dislikes: 290
 * Total Accepted:    298.7K
 * Total Submissions: 917.2K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 *
 * Note:
 *
 * The solution set must not contain duplicate quadruplets.
 *
 * Example:
 *
 *
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 *
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */
package leetcode18

// @lc code=start
import "sort"

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target { // 最小的可能
			break
		}
		if nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target { // 最大的可能
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target { // 最小的可能
				break
			}
			if nums[i]+nums[j]+nums[length-1]+nums[length-2] < target { // 最大的可能
				continue
			}
			ps := nums[i] + nums[j]
			l := j + 1
			r := length - 1
			for l < r {
				if s := ps + nums[l] + nums[r]; s > target { // 和大于 0，则右边向前移一位
					r--
				} else if s < target { // 和小于 0，则左边向后移一位
					l++
				} else { // 同时移位
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
					l++
				}
			}
		}
	}
	return res
}

// @lc code=end
