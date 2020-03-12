/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (31.76%)
 * Likes:    2786
 * Dislikes: 964
 * Total Accepted:    324.6K
 * Total Submissions: 1M
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 *
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 *
 * The replacement must be in-place and use only constant extra memory.
 *
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 *
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */
package leetcode31

// @lc code=start
func nextPermutation(nums []int) {
	var i int
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] { // 找到最靠右的小数
			for k := len(nums) - 1; k >= i; k-- {
				if nums[i-1] < nums[k] { // 因为靠右的是降序，从右开始找第一个比小数大的数，就是最小的大数
					nums[i-1], nums[k] = nums[k], nums[i-1]
					break
				}
			}
			break
		}
	}
	// reverse
	for x, y := i, len(nums)-1; x < y; x, y = x+1, y-1 {
		nums[x], nums[y] = nums[y], nums[x]
	}
}

// @lc code=end
