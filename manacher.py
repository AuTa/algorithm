import pytest 


class Manacher:

    def __init__(self, string):
        self.o_string = string
        self.string = self.handle_string()
    
    def handle_string(self):
        return '$#{}#&'.format('#'.join(self.o_string)) 

    def max_len(self):
        id_ = mx = max_len = 0
        p = {}
        for i in range(1, len(self.string) - 1):
            if i < mx:
                p[i] = min(p[2 * id_ - i], mx - i)
            else:
                p[i] = 1
            while self.string[i - p[i]] == self.string[i + p[i]]:
                p[i] += 1
            if mx < i + p[i]:
                mx = i + p[i]
                id_ = i
            max_len = max(max_len, p[i] - 1)
        return max_len

    def longest_palindrome(self):
        id_ = mx = max_len = 0
        palindrome = ''
        p = {}
        for i in range(1, len(self.string) - 1):
            if i < mx:
                p[i] = min(p[2 * id_ - i], mx - i)
            else:
                p[i] = 1
            while self.string[i - p[i]] == self.string[i + p[i]]:
                p[i] += 1
            if mx < i + p[i]:
                mx = i + p[i]
                id_ = i
            if p[i] - 1 > max_len:
                palindrome = self.string[i - p[i] + 1:i + p[i]].replace('#', '')
                max_len = p[i] - 1
        return palindrome


if __name__ == '__main__':
    s = Manacher('babad')
    print(s.max_len())
    print(s.longest_palindrome())
