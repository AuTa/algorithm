# [848. Shifting Letters](https://leetcode.com/problems/shifting-letters/)

We have a string `S` of lowercase letters, and an integer array `shifts`.

Call the shift of a letter, the next letter in the alphabet, (wrapping around so that `'z'` becomes `'a'`). 

For example, `shift('a') = 'b'`, `shift('t') = 'u'`, and `shift('z') = 'a'`.

Now for each `shifts[i] = x`, we want to shift the first `i+1` letters of `S`, `x` times.

Return the final string after all such shifts to `S` are applied.

### Example 1:
```md
Input: S = "abc", shifts = [3,5,9]
Output: "rpl"
Explanation: 
We start with "abc".
After shifting the first 1 letters of S by 3, we have "dbc".
After shifting the first 2 letters of S by 5, we have "igc".
After shifting the first 3 letters of S by 9, we have "rpl", the answer.
```
### Note:

1. `1 <= S.length = shifts.length <= 20000`
2. `0 <= shifts[i] <= 10 ^ 9`

## 思路

### 第一版

遍历 shifts，索引为 i，然后遍历 i 次 S，shift 每个字符，如果结果大于 'z'，将结果减去 26。 
这一版太慢了。

### 第二版

第一个字符的 shift 的位数是 shifts 中所有数字的和，如果反响遍历 shifts，保存已经遍历的数字和，这个和就是当前位的 S 的字符需要 shift 的位数。  
求和应该取 26 的余数（不然和超过 uint8 的范围导致相加失败）。  
新的值等于 (旧值 + shift 的位数 - 'a') 取 26 的余数 + 'a'。通过相对于 'a' 的偏移减少是否超过 'z' 的判断。
