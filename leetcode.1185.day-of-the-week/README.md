# [Day of the Week](https://leetcode.com/problems/day-of-the-week/description/)

## 基姆拉尔森计算公式（Kim larsson calculation formula）

w = (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400 + 1) % 7

0 表示周日。把一月和二月看成是上一年的十三月和十四月。

[基姆拉尔森计算公式 推导](https://www.cnblogs.com/SeekHit/p/7498408.html)

## 天数累计模 7

1971/01/01 是周五，那么第零天的累计量是 4。  
累计年份天数，再累计月份天数，再加上最后一个月天数，最后加上偏移量 4。
