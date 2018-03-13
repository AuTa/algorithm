# [668. Kth Smallest Number in Multiplication Table](https://leetcode.com/problems/kth-smallest-number-in-multiplication-table)

## Description

Nearly every one have used the Multiplication Table. But could you find out the k-th smallest number quickly from the multiplication table?

Given the height m and the length n of a m * n Multiplication Table, and a positive integer k, you need to return the k-th smallest number in this table.

### Example 1:
> **Input:** m = 3, n = 3, k = 5  
**Output:**   
**Explanation:**    
The Multiplication Table:  
1	2	3  
2	4	6  
3	6	9  
> 
> The 5-th smallest number is 3 (1, 2, 2, 3, 3).

### Example 2:
> **Input:** m = 2, n = 3, k = 6  
**Output:**   
**Explanation:**   
The Multiplication Table:  
1	2	3  
2	4	6  
>
> The 6-th smallest number is 6 (1, 2, 2, 3, 4, 6).

### Note:
1. The m and n will be in the range [1, 30000].
1. The k will be in the range [1, m * n]

## 思路

### Binary search 二分搜索

- 计算 `[low, high]` 中间数 x，不大于中间数的数共有在 `x_c` 个
- `x_c < k` 则 `k` 在 `[x + 1, high]` 之间
- `x_c >= k` 则 `k` 在 `[low, x]` 之间（`x` 可能不是乘法表里的数）
- 因为 `low < high`，对于 `[x + 1, high]` 如果 `x + 1 == high` 则第 `k` 个数是 `x + 1`，对于 `[low, x]` 如果 `low == x` 则第 `k` 个数是 `low`

计算不大于 x 的数有多少个的方法(每一列有多少个小于 x 的数)：
- 第 1 列 min(x // 1, n)
- 第 2 列 min(x // 2, n)
- ...
- 第 i 列 min(x // i, n)，i <= m
