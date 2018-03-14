# http://www.lintcode.com/zh-cn/problem/matrix-zigzag-traversal/


class Solution:
    """
    @param matrix: An array of integers
    @return: An array of integers
    """
    def printZMatrix(self, matrix):
        # write your code here
        """遇到边界后重新计算 next x y"""
        m, n = len(matrix[0]), len(matrix)
        res = []
        x = y = 0
        direct = [1, -1]
        while (x + 1) * (y + 1) <= m * n:
            res.append(matrix[y][x])
            next_x = x + direct[0]
            next_y = y + direct[1]
            if (next_x < 0 or next_x >= m) and (next_y < 0 or next_y >= n):
                if direct == [-1, 1]:
                    next_y = y
                    next_x = x + 1
                else:
                    next_x = x
                    next_y = y + 1
                direct = direct[::-1]
            elif next_x < 0 or next_x >= m:
                next_x = x
                next_y = y + 1
                direct = direct[::-1]
            elif next_y < 0 or next_y >= n:
                next_y = y
                next_x = x + 1
                direct = direct[::-1]
            x, y = next_x, next_y
        return res
