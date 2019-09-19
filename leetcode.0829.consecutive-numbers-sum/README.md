# [Consecutive Numbers Sum](https://leetcode.com/problems/consecutive-numbers-sum/)

Given a positive integer `N`, how many ways can we write it as a sum of consecutive positive integers?

**Example 1:**
```md
**Input:** 5
**Output:** 2
**Explanation:** 5 = 5 = 2 + 3
```
**Example 2:**
```md
**Input:** 9
**Output:** 3
**Explanation:** 9 = 9 = 4 + 5 = 2 + 3 + 4
```
**Example 3:**
```md
**Input:** 15
**Output:** 4
**Explanation:** 15 = 15 = 8 + 7 = 4 + 5 + 6 = 1 + 2 + 3 + 4 + 5
```
**Note:** `1 <= N <= 10 ^ 9`.

## 思路

### 解法一
如果可以是 `i` 个连续数字相加，那么就可以是 `(1 + ... + i) + i * x = N`，也就是 `N - (1 + ... + i)` 可以被 `i` 整除。

### 解法二
`i = b - a + 1`，`((a + b) * i) / 2 = N` 得到 `a + b = 2N / i`，即 `2N` 可以被 `i` 整除，然后解二元方程，得到 `2b = 2N / i + i - 1` 和 `2a = 2N / i - i + 1`，即 `2N / i + i - 1` 可以被 `2` 整除，`2N / i - i + 1` 可以被 `2` 整除。
