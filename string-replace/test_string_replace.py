from string_replace import Solution


def test_string_replace():
    solution = Solution()
    a = ["ab","aba"]
    b = ["cc","ccc"]
    s = "ababa"
    assert('cccba', solution.stringReplace(a, b, s))
