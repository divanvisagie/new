import pytest

from template import read, enter_loop_with

def test_given_template_user_is_prompted_twice():
    called_count = 0
    def prompt_mock(p):
        nonlocal called_count 
        called_count += 1

    test_object = {
        'replace': {
            'strings': [
                {
                    'match': 'com.divanvisagie.example', 
                    'description': 'The package name'
                }, 
                {
                    'match': 'Hello World',
                    'description': 'The string that is printed by the application'
                }
            ]
        }
    }

    enter_loop_with(test_object, prompt=prompt_mock)
    assert called_count == 2

def test_given_template_map_is_returned():
    called_count = 0
    def prompt_mock(p):
        nonlocal called_count 
        called_count += 1
        if called_count == 1:
            return 'com.example.test'
        else:
            return 'Hello Test'

    test_object = {
        'replace': {
            'strings': [
                {
                    'match': 'com.divanvisagie.example', 
                    'description': 'The package name'
                }, 
                {
                    'match': 'Hello World',
                    'description': 'The string that is printed by the application'
                }
            ]
        }
    }

    actual = enter_loop_with(test_object, prompt=prompt_mock)
    assert actual[0] == ('com.divanvisagie.example', 'com.example.test')
    assert actual[1] == ('Hello World', 'Hello Test')
