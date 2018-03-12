# [697. Degree of an Array](https://leetcode.com/problems/degree-of-an-array/description/)

## 描述

Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

### Example 1:
> **Input:** [1, 2, 2, 3, 1]  
> **Output:** 2  
> **Explanation:**   
> The input array has a degree of 2 because both elements 1 and 2 appear twice.  
> Of the subarrays that have the same degree:  
> [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]  
> The shortest length is 2. So return 2.
### Example 2:
> **Input:** [1,2,2,3,1,4,2]  
> **Output:** 6

## 思路

统计每个数字出现次数和初次出现索引，如果次数等于原 degree 最小长度是之前最小长度和当前长度的较小值，如果次数大于原 degree 则更新 degree 和最小长度。
