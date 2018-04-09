class Solution:
    def removeComments(self, source):
        """
        :type source: List[str]
        :rtype: List[str]
        """
        result, buffer, block_start = [], '', False  # 结果, 缓冲, 块注释开始标记
        for s in source:
            i = 0
            while i < len(s):
                if block_start is False:
                    if s[i:i+2] == '//':  # 行注释
                        i = len(s)
                    elif s[i:i+2] == '/*':  # 块注释开始
                        block_start = True
                        i += 2
                    else:
                        buffer += s[i]
                        i += 1
                else:
                    if s[i:i+2] == '*/':
                        block_start = False
                        i += 2
                    else:
                        i += 1
            if buffer and block_start is False:
                result.append(buffer)
                buffer = ''
        return result
