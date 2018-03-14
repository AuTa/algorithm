class Solution:
    def combinationSum(self, candidates, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        candidates.sort()
        res = []
        def dfs(remain, path):
            if remain == 0:
                res.append(path)
                return
            for num in candidates:
                if num > remain:
                    break
                if path and num < path[-1]: 
                    continue
                else:
                    dfs(remain - num, path + [num])
        dfs(target, [])
        return res
