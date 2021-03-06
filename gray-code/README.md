# [411. 格雷编码](http://www.lintcode.com/zh-cn/problem/gray-code/)

## 描述

格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个二进制的差异。  
给定一个非负整数 `n` ，表示该代码中所有二进制的总数，请找出其格雷编码顺序。一个格雷编码顺序必须以 `0` 开始，并覆盖所有的 `2n` 个整数。

> ### 注意事项  
> 对于给定的 `n`，其格雷编码顺序并不唯一。  
> 根据以上定义， `[0,2,3,1]` 也是一个有效的格雷编码顺序。

### 样例 
给定 `n = 2`， 返回 `[0,1,3,2]`。其格雷编码顺序为：
```
00 - 0
01 - 1
11 - 3
10 - 2
```

## 思路

```
f(0) = [0]
f(1) = [0, 1] = [0] + [0+1] = f(0) + [x+1 for x in f(0)]
f(2) = [0, 1, 3, 2] = [0, 1] + [1 + 2, 0 + 2] = f(1) + [x+2**1 for x in f(1)[::-1]]
...
f(n) = f(n-1) + [2 ** (n - 1) + x for x in f(n-1)[::-1]]
```
