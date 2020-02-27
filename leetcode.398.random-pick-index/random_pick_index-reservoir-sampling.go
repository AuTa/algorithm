/*
 * @lc app=leetcode id=398 lang=golang
 *
 * [398] Random Pick Index
 *
 * https://leetcode.com/problems/random-pick-index/description/
 *
 * algorithms
 * Medium (53.02%)
 * Likes:    354
 * Dislikes: 564
 * Total Accepted:    70.4K
 * Total Submissions: 132.7K
 * Testcase Example:  '["Solution","pick"]\n[[[1,2,3,3,3]],[3]]'
 *
 * Given an array of integers with possible duplicates, randomly output the
 * index of a given target number. You can assume that the given target number
 * must exist in the array.
 *
 * Note:
 * The array size can be very large. Solution that uses too much extra space
 * will not pass the judge.
 *
 * Example:
 *
 *
 * int[] nums = new int[] {1,2,3,3,3};
 * Solution solution = new Solution(nums);
 *
 * // pick(3) should return either index 2, 3, or 4 randomly. Each index should
 * have equal probability of returning.
 * solution.pick(3);
 *
 * // pick(1) should return 0. Since in the array only nums[0] is equal to 1.
 * solution.pick(1);
 *
 *
 */
package leetcode0398

// @lc code=start
import (
	"math/rand"
)

type Solution struct {
	n []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	count := 0
	res := 0
	for i, v := range this.n {
		if v == target {
			count++
			if rand.Intn(count) == 0 {
				res = i
			}
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end
