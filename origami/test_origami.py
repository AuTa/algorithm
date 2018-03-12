from origami import Solution


def test_origami():
    solution = Solution()
    assert('0010011', solution.getString(3))
