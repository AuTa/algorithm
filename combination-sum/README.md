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
