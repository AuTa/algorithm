class Solution:
    def findKthNumber(self, m, n, k):
        """
        :type m: int
        :type n: int
        :type k: int
        :rtype: int
        """
        low = 1
        high = m * n
        
        while low < high:
            middle = (low + high) // 2
            middle_place = sum(min(middle // i, n) for i in range(1, min(middle + 1, m + 1)))
            if middle_place >= k:
                high = middle
            else:
                low = middle + 1
        return low
