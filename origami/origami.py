class Solution:
    """
    @param n: The folding times
    @return: the 01 string
    """
    def getString(self, n):
        # Write your code here
        res = ''
        i = 0
        def next_num():
            # 依次返回 0、1
            while True:
                yield '0'
                yield '1'
        while i < n:
            n_num = next_num()
            new_res = next(n_num)
            for x in res:
                new_res += x + next(n_num)
            res = new_res
            i += 1
        return res
