/*
 * @lc app=leetcode id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 *
 * https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (27.84%)
 * Likes:    270
 * Dislikes: 45
 * Total Accepted:    10.9K
 * Total Submissions: 39.1K
 * Testcase Example:  '13\n2'
 *
 * Given integers n and k, find the lexicographically k-th smallest integer in
 * the range from 1 to n.
 *
 * Note: 1 ≤ k ≤ n ≤ 10^9.
 *
 * Example:
 *
 * Input:
 * n: 13   k: 2
 *
 * Output:
 * 10
 *
 * Explanation:
 * The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so
 * the second smallest number is 10.
 *
 *
 *
 */
package leetcode440

// @lc code=start
func findKthNumber(n int, k int) int {
	// 十叉树，最初当前节点是 1
	currentNode := 1
	// 第一个节点占用了一个位置，所以找 k-1 位置的元素
	k--
	for k > 0 {
		firstNode := currentNode    // 子树的第一个节点
		lastNode := currentNode + 1 // 下一个子树的第一个节点，用来计算增加的节点数
		currentTreeNodes := 0       // 当前这个子树有多少节点
		// 深度优先，先序遍历直到最大值的树
		for firstNode <= n {
			// 因为这棵树不一定是满十叉树，所以增加的节点取较小值，n+1 是因为多减了一个元素
			if lastNode < n+1 {
				currentTreeNodes += lastNode - firstNode
			} else {
				currentTreeNodes += n + 1 - firstNode
			}
			// 进入下一层
			firstNode *= 10
			lastNode *= 10
		}
		if currentTreeNodes > k { // 子树的子节点比 k 大，说明结果在这棵树中，向下递归寻找
			k--
			currentNode *= 10
		} else { // 不在树中，则到下一棵树中找第 k-currentTreeNodes 位的元素
			k -= currentTreeNodes
			currentNode++
		}
	}
	return currentNode
}

// @lc code=end
