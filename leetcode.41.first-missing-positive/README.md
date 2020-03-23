# [First Missing Positive](https://leetcode.com/problems/first-missing-positive/description/) - 缺失的第一个正数

将每个数放到它应该待的索引位置，然后遍历处理后的数组，如果索引位置和值不符合则返回这个索引。

因为题目要求的是正数，那么 `1` 应该处在索引为 `0` 处，即每个值的索引应该 `-1`。非正数不用考虑、大于当前数组长度的数不用考虑。因为可能有重复的数，所以替换时要避免替换的两个数相同，导致无限循环。
