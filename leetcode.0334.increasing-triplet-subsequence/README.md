# [334. Increasing Triplet Subsequence](https://leetcode.com/problems/increasing-triplet-subsequence/)

Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.

Formally the function should:

> Return true if there exists *i, j, k*  
> such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.

**Note:** Your algorithm should run in O(n) time complexity and O(1) space complexity.

### Example 1:
```
Input: [1,2,3,4,5]
Output: true 
```
### Example 2:
```
Input: [5,4,3,2,1]
Output: false
```

## 思路

保持大值和小值接近，之后有一个数大于大值时就符合条件。新值小于小值则更新小值，新值小于大值则更新大值。
