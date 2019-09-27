# [Distribute Coins in Binary Tree](https://leetcode.com/problems/distribute-coins-in-binary-tree/)

## 思路

**DFS** 深度优先搜索

每个节点只能保留一个硬币，不管是移入还是移出，都需要移动`当前硬币 - 1` 次，从最深处开始移动，每次移动的数都累计在父节点上，父节点需要移动的次数是`自身硬币数 + 左边移动的硬币数 + 右边移动的硬币数 - 1`。
