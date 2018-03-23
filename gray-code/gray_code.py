class Solution:
    """
    @param n: a number
    @return: Gray code
    """
    def grayCode(self, n):
        # write your code here
        result = [0]
        i = 0
        while i < n:
            prefix = 1 << i
            result.extend([(prefix + x) for x in result[::-1]])
            i += 1

        return result
