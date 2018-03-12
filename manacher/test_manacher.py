import pytest
from manacher import Manacher

@pytest.fixture()
def babad():
    return Manacher('babad')

def test_babad_max_len(babad):
    assert babad.max_len() == 3

def test_babad_longest_palindrome(babad):
    assert babad.longest_palindrome() == 'bab'

@pytest.fixture()
def abbahopxpo():
    return Manacher('abbahopxpo')

def test_abbahopxpo_max_len(abbahopxpo):
    assert abbahopxpo.max_len() == 5

def test_abbahopxpo_longest_palindrome(abbahopxpo):
    assert abbahopxpo.longest_palindrome() == 'opxpo'
