class Solution:
    """
    @param a: The A array
    @param b: The B array
    @param s: The S string
    @return: The answer
    """
    def stringReplace(self, a, b, s):
        # Write your code here
        a_lens = sorted({len(x) for x in a}, reverse=True)
        a_index = {x: i for i, x in enumerate(a)}  # 通过 hash 加快查找速度
        answer = ''
        while s:
            finded = False
            for length in a_lens:
                if length > len(s):
                    continue
                if s[:length] in a_index:
                    answer += b[a_index[s[:length]]]
                    s = s[length:]
                    finded = True
                    break
            if finded is False:
                answer += s[0]
                s = s[1:]
        return answer
