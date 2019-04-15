# [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)

Given two strings s and t , write a function to determine if t is an anagram of s.

### Example 1:
```md
Input: s = "anagram", t = "nagaram"
Output: true
```
### Example 2:
```md
Input: s = "rat", t = "car"
Output: false
```
### Note:
You may assume the string contains only lowercase alphabets.

### Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?

## 思路

一个字符串对字符做递增，另一个字符串对字符做递减，最后结果如果所有字符都为 0，则符合验证。
