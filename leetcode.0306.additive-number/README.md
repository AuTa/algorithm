# [306. Additive Number](https://leetcode.com/problems/additive-number/)

Additive number is a string whose digits can form additive sequence.

A valid additive sequence should contain **at least** three numbers. Except for the first two numbers, each subsequent number in the sequence must be the sum of the preceding two.

Given a string containing only digits `'0'-'9'`, write a function to determine if it's an additive number.

**Note:** Numbers in the additive sequence **cannot** have leading zeros, so sequence `1, 2, 03` or `1, 02, 3` is invalid.

**Example 1:**
```md
**Input:** "112358"
**Output:** true 
**Explanation:** The digits can form an additive sequence: 1, 1, 2, 3, 5, 8. 
             1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
```
**Example 2:**
```md
**Input:** "199100199"
**Output:** true 
**Explanation:** The additive sequence is: 1, 99, 100, 199. 
             1 + 99 = 100, 99 + 100 = 199
```
**Follow up:**  
How would you handle overflow for very large input integers?

## 思路

从头开始选第一个数、第二个数和剩下的数，判断剩下的数开头是不是第一个数和第二个数的和，如果是则第一个数换成第二个数，第二个数换成和，剩下的数排除为和的数，再判断剩下的数，直到没有剩下的数，返回 `true`，否则第二个数加一位，然后继续判断，如果第二个数取到了最后一位还没有返回 `true`，那么第一个数加一位，继续下去。

如果第一个数是 `0` 开头并且不等于 `0` 那么可以直接返回 `false`。  
如果第二个数是 `0` 开头并且不等于 `0` 那么直接递增第一个数。
