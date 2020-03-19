# [39. Combination Sum](https://leetcode.com/problems/combination-sum/)

## Description

Given a set of candidate numbers (C) (without duplicates) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

### Note:

- All numbers (including target) will be positive integers.
- The solution set must not contain duplicate combinations.

For example, given candidate set `[2, 3, 6, 7]` and target `7`,   
A solution set is:   
```
[
  [7],
  [2, 2, 3]
]
```

## 思路

DFS 深度优先搜索

循环 candidates，计算每一个数还差多少到 target，差的数作为新的 target 重复执行，直到剩下到数等于 0。

## 回溯

回溯算法框架。解决一个回溯问题，实际上就是一个决策树的遍历过程。你只需要思考 3 个问题：

1. 路径：也就是已经做出的选择。
2. 选择列表：也就是你当前可以做的选择。
3. 结束条件：也就是到达决策树底层，无法再做选择的条件。

代码方面，回溯算法的框架：
```
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

for 选择 in 选择列表:
    做选择
    backtrack(路径, 选择列表)
    撤销选择
```
