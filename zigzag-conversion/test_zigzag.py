import pytest
from zigzag_conversion import ZigZagConversion

def test_convert():
    zig_zag = ZigZagConversion()
    assert zig_zag.convert('PAYPALISHIRING', 3) == 'PAHNAPLSIIGYIR'
