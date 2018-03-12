class Solution:
    def findShortestSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        degree = 1
        degree_nums = {nums[0]: [1, 0]}  # 暂存数字出现次数和初次出现索引
        result = 1
        for i in range(1, len(nums)):
            if nums[i] in degree_nums:
                degree_nums[nums[i]][0] += 1
                i_degree = degree_nums[nums[i]][0]
                if i_degree == degree:
                    result = min(result, i - degree_nums[nums[i]][1] + 1)
                    degree = i_degree
                elif i_degree > degree:
                    result = i - degree_nums[nums[i]][1] + 1
                    degree = i_degree
            else:
                i_degree = 1
                degree_nums[nums[i]] = [1, i]
        return result
