# [Next Permutation](https://leetcode.com/problems/next-permutation/description/)

所谓下一个排列就是当前数字列表中数字组成的列表中比当前列表大的最小的列表。

为了得到这样的列表，分析如下：
1. 下一个数比当前数大，那么只需要将当前数中的一个比较小的数替换成这个小数后面的大数；
2. 下一个数要满足增长的幅度尽可能小，那么小数应该是尽量靠右的数，被交换的大数应该要尽可能小，后面的数应该改为从小到大排列。

过程：
1. 从后开始往前找第一个 `nums[i-1] < nums[i]`，`nums[i, end]` 必然是降序;
2. 再从后开始往前找第一个 `nums[i-1] < nums[k]`，交换大数和小数；
3. 因为 `nums[i-1] >= nums[k+1]`，所以交换后的 `nums[i, end]` 还是降序，反转为升序。
4. 如果本身就是降序，那么 1 的过程中找不到匹配的元素，`i = 0`，同样可以使用 3 的过程反转。