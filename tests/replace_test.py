import pytest

from new.replace import replace_strings, get_file_paths, replace_match

def test_get_file_paths():
    actual = get_file_paths('./tests/example_tree')
    assert './tests/example_tree/test.txt' in actual
    assert './tests/example_tree/onedeep/other_file.txt' in actual

def test_replace_match():
    test_content = 'We are the cheese cake factory'
    test_tuple = ('We','You')
    actual = replace_match(test_tuple, test_content)
    assert actual == 'You are the cheese cake factory'

    