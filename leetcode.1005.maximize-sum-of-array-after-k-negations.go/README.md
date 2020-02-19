# [Maximize Sum Of Array After K Negations](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/)

*贪心算法*

负数占用一次转换。然后如果剩下最小的是 0，则可以转换任何多次。如果剩下的次数是偶数，则可以将同一个数字转换两次保持不变；如果剩下的次数是奇数，则将现在最小的数转换为负数则可以。判断现在最小的数可以用当前数字和前一个又负数转换成的数比较大小。
