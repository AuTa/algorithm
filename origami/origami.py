class Solution:
    """
    @param n: The folding times
    @return: the 01 string
    """
    def getString(self, n):
        # Write your code here
        res = '0'
        if n == 1:
            return res
        i = 1
        while i < n:
            res = res + '{r:0{l}b}'.format(l=len(res) + 1, r=int(res[::-1], 2) ^ ((1 << len(res)) - 1))
            i += 1
        return res
