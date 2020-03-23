/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (30.73%)
 * Likes:    2759
 * Dislikes: 730
 * Total Accepted:    293.6K
 * Total Submissions: 949.7K
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array, find the smallest missing positive
 * integer.
 *
 * Example 1:
 *
 *
 * Input: [1,2,0]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: [3,4,-1,1]
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: [7,8,9,11,12]
 * Output: 1
 *
 *
 * Note:
 *
 * Your algorithm should run in O(n) time and uses constant extra space.
 *
 */
package leetcode41

// @lc code=start
func firstMissingPositive(nums []int) int {
	l := len(nums)
	for i := range nums {
		for nums[i] != i+1 && nums[i] > 0 && nums[i] < l && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < l; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return l + 1
}

// @lc code=end
