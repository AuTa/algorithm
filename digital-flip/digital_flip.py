class Solution:
    """
    @param nums: the array
    @return: the minimum times to flip digit
    """
    def flipDigit(self, nums):
        # Write your code here
        # 最坏的情况是所有 1 翻成 0
        # 如果碰到 1 则前面的 0 都翻成 1，剩下的 1 都翻成 0，或者之前的最坏情况
        one_count = sum(nums)
        ans = one_count
        temp = 0  # 出现的 1 的个数
        for i, num in enumerate(nums):
            if num == 1:
                zore_count = i - temp
                temp += 1
                rest_one_count = one_count - temp
                ans = min(ans, zore_count + rest_one_count)
        return ans

    def flipDigit2(self, nums):
        # Write your code here
        counts = [0, 0]  # 翻成 0，1 的次数
        if nums[-1] == 1:
            counts[0] += 1
        else:
            counts[1] += 1
        for i in range(len(nums) - 2, -1, -1):
            if nums[i] == 1:
                counts[1] = min(counts)
                counts[0] += 1
            else:
                counts[1] = min(counts) + 1
        return min(counts)
