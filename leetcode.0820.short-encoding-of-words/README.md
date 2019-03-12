# [820. Short Encoding of Words](https://leetcode.com/problems/short-encoding-of-words/)

## Description

Given a list of words, we may encode it by writing a reference string S and a list of indexes A.

For example, if the list of words is ["time", "me", "bell"], we can write it as S = "time#bell#" and indexes = [0, 2, 5].

Then for each index, we will recover the word by reading from the reference string from that index until we reach a "#" character.

What is the length of the shortest reference string S possible that encodes the given words?

### Example:

> Input: words = ["time", "me", "bell"]  
> Output: 10  
> Explanation: S = "time#bell#" and indexes = [0, 2, 5].
 

### Note:

> 1 <= words.length <= 2000.  
> 1 <= words[i].length <= 7.  
> Each word has only lowercase letters.

## 思路

1. 列表按由长到短排序后，依次判断每个单词加上 # 是否在结果字符串上，不在则添加在结果字符串后。
2. 将列表放入 hashset，从 hashset 中移除列表中每个单词的每个带尾部的子字符串。

第一种方案是邪道，特别耗时。
