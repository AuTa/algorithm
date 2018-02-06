class ZigZagConversion:

    def convert(self, s, num_rows):
        """https://leetcode.com/problems/zigzag-conversion/description/
        
        """
        if num_rows == 1 or num_rows >= len(s):
            return s
        zigzag_list = [''] * num_rows
        index, direction = 0, 1
        max_index = num_rows - 1
        for i in s:
            zigzag_list[index] += i
            if index == 0:
                direction = 1
            elif index == max_index:
                direction = -1
            index += direction
        return ''.join(zigzag_list)
