# 239. [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)

Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

### Example:
```
Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7] 
Explanation: 

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
```
### Note:
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.

### Follow up:
Could you solve it in linear time?

## 思路

用队列保存窗口内最大值以及最大值之后的待定值；  
每次添加新值时从右向左移除比新值小的值，因为这些待定值永远用不上了；  
如果当前索引减去窗口大小等于队列最左索引，即最左索引需要被移除；  
此时最左索引为当前窗口最大值。

## TODO

神奇的解法：  
[ ] [O(n) solution in Java with two simple pass in the array](https://leetcode.com/problems/sliding-window-maximum/discuss/65881/O(n)-solution-in-Java-with-two-simple-pass-in-the-array)
