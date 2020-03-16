/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (33.59%)
 * Likes:    3892
 * Dislikes: 407
 * Total Accepted:    590K
 * Total Submissions: 1.8M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 *
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 *
 * You may assume no duplicate exists in the array.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * Example 1:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 *
 */
package leetcode33

// @lc code=start
func search(nums []int, target int) int {
	var mid int
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid = (low + high) / 2
		if target == nums[mid] {
			return mid
		}
		if !((nums[low] <= target) != (target <= nums[mid]) != (nums[mid] < nums[low])) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// @lc code=end
