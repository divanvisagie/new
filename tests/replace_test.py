import pytest

from new.replace import name_in_line, replace_match, get_file_paths, name_in_line

def test_get_file_paths():
    actual = get_file_paths('./tests/example_tree')
    assert './tests/example_tree/test.txt' in actual
    assert './tests/example_tree/onedeep/other_file.txt' in actual

def test_replace_match():
    test_content = 'We are the cheese cake factory'
    test_tuple = ('We','You')
    actual = replace_match(test_tuple, test_content)
    assert actual == 'You are the cheese cake factory'


def test_find_similar_names_should_return_lines():
    original_project_name = 'my-fake-template'
    test_data = [
        { 
            'input': 'import com.test.myfaketemplate',
            'output': 'com.test.myfaketemplate'
        },
        {
            'input': 'my-fake-template',
            'output': 'my-fake-template'
        },
        {
            'input': 'class MyFakeTemplate extends Chicken',
            'output': 'MyFakeTemplate'
        }
    ]
    for item in test_data:
        actual = name_in_line(item['input'], original_project_name)
        assert actual == item['output']

