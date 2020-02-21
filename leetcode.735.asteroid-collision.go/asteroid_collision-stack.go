/*
 * @lc app=leetcode id=735 lang=golang
 *
 * [735] Asteroid Collision
 *
 * https://leetcode.com/problems/asteroid-collision/description/
 *
 * algorithms
 * Medium (39.94%)
 * Likes:    713
 * Dislikes: 82
 * Total Accepted:    44.2K
 * Total Submissions: 110.7K
 * Testcase Example:  '[5,10,-5]'
 *
 *
 * We are given an array asteroids of integers representing asteroids in a
 * row.
 *
 * For each asteroid, the absolute value represents its size, and the sign
 * represents its direction (positive meaning right, negative meaning left).
 * Each asteroid moves at the same speed.
 *
 * Find out the state of the asteroids after all collisions.  If two asteroids
 * meet, the smaller one will explode.  If both are the same size, both will
 * explode.  Two asteroids moving in the same direction will never meet.
 *
 *
 * Example 1:
 *
 * Input:
 * asteroids = [5, 10, -5]
 * Output: [5, 10]
 * Explanation:
 * The 10 and -5 collide resulting in 10.  The 5 and 10 never collide.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * asteroids = [8, -8]
 * Output: []
 * Explanation:
 * The 8 and -8 collide exploding each other.
 *
 *
 *
 * Example 3:
 *
 * Input:
 * asteroids = [10, 2, -5]
 * Output: [10]
 * Explanation:
 * The 2 and -5 collide resulting in -5.  The 10 and -5 collide resulting in
 * 10.
 *
 *
 *
 * Example 4:
 *
 * Input:
 * asteroids = [-2, -1, 1, 2]
 * Output: [-2, -1, 1, 2]
 * Explanation:
 * The -2 and -1 are moving left, while the 1 and 2 are moving right.
 * Asteroids moving the same direction never meet, so no asteroids will meet
 * each other.
 *
 *
 *
 * Note:
 * The length of asteroids will be at most 10000.
 * Each asteroid will be a non-zero integer in the range [-1000, 1000]..
 *
 */
package leetcode735

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	ans := make([]int, 0, len(asteroids))
	for _, a := range asteroids {
		for {
			if l := len(ans); l == 0 || a >= 0 || ans[l-1] < 0 {
				// 栈内没有行星、或新行星向右行驶、或栈顶行星向左行驶则将新行星压入栈顶
				ans = append(ans, a)
				break
			} else {
				if ans[l-1] < -a {
					// 栈顶行星比新行星小，则栈顶行星爆炸，移除后继续和栈剩下的判断
					ans = ans[:l-1]
					continue
				} else if ans[l-1] == -a {
					// 栈顶行星和新行星一样大，同时爆炸，跳出判断
					ans = ans[:l-1]
				} // 栈顶行星比新行星大，新行星爆炸，跳出判断
				break
			}
		}
	}
	return ans
}

// @lc code=end
