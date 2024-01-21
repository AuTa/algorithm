# [60. Permutation Sequence](https://leetcode.com/problems/permutation-sequence/)

## Description

The set `[1,2,3,...,n]` contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

1. "123"
2. "132"
3. "213"
4. "231"
5. "312"
6. "321"
   
Given n and k, return the kth permutation sequence.

### Note:

- Given n will be between 1 and 9 inclusive.
- Given k will be between 1 and n! inclusive.
  
### Example 1:

> Input: n = 3, k = 3  
> Output: "213"
### Example 2:

> Input: n = 4, k = 9  
> Output: "2314"

## 思路

每一次循环都是 (n-1)! 个数，当前第一个数即 k 包含多少的完整循环，然后从切片索引中取出，下一个 k 是完整循环之后的余数，然后递归查找下一个数。

[康托展开](https://zh.wikipedia.org/wiki/康托展开)
