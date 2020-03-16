# [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array)

思路使用二分搜索。

当低位比中间位小时，说明低到中是升序，如果 target 也处于这个区间内，则下次搜索换到低位区间。`nums[low] <= target <= nums[mid]`  
当低位比中间为大时，说明中到高是升序，如果 target 比中间位大或者 target 比低位小，则下次搜索换到高位区间。`target <= nums[mid] < nums[low] || nums[mid] < nums[low] <= target`

上述情况可以归纳成：
```
nums[low] <= target <= nums[mid]
             target <= nums[mid] < nums[low]
                       nums[mid] < nums[low] <= target
```

如果为真满足如下条件：
- `nums[low] <= target` && `target <= nums[mid]` 
- `target <= nums[mid]` && `nums[mid] < nums[low]`
- `nums[mid] < nums[low]` && `nums[low] <= target`

即 `nums[low] <= target`、`target <= nums[mid]`、 `nums[mid] < nums[low]` 三者任意两个为真，另一个为假。  
使用异或操作可以得到上述结果（两项为真时异或结果为假，一项为真时异或结果为真）。

golang 中 boolean 的异或用 `A != B` 
