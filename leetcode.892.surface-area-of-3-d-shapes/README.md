# [Surface Area of 3D Shapes](https://leetcode.com/problems/surface-area-of-3d-shapes/description/)

总面积减去重叠的面积。

其中：

- 同一个位置重叠面积是 `2*(个数-1)`
- 相邻位置重叠面积是 `2*min((i,j)个数, (i-1,j)个数)` 和 `2*min((i,j)个数, (i,j-1)个数)`
