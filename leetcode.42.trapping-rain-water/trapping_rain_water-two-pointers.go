/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (46.72%)
 * Likes:    5869
 * Dislikes: 107
 * Total Accepted:    445.4K
 * Total Submissions: 945.9K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it is able to trap after raining.
 *
 *
 * The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
 * In this case, 6 units of rain water (blue section) are being trapped. Thanks
 * Marcos for contributing this image!
 *
 * Example:
 *
 *
 * Input: [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 *
 */
package leetcode42

// @lc code=start
func trap(height []int) int {
	var ans int
	var left, right = 0, len(height) - 1
	var left_max, right_max int
	for left < right {
		if height[left] < height[right] {
			if height[left] >= left_max {
				left_max = height[left]
			} else { // 左边比它高，右边也比它高
				ans += left_max - height[left]
			}
			left++
		} else {
			if height[right] >= right_max {
				right_max = height[right]
			} else {
				ans += right_max - height[right]
			}
			right--
		}
	}
	return ans
}

// @lc code=end

func trap_stack(height []int) int {
	var ans int
	stack := make([]int, 0)
	l := len(stack)
	for cur := 0; cur < len(height); cur++ {
		for l != 0 && height[cur] > height[stack[l-1]] {
			top := stack[l-1]
			stack = stack[:l-1]
			l--
			if l == 0 {
				break
			}
			distance := cur - stack[l-1] - 1
			bounded_height := min(height[cur], height[stack[l-1]]) - height[top]
			ans += distance * bounded_height
		}
		stack = append(stack, cur)
		l++
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
