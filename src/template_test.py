import pytest

from template import read, enter_loop_with

def test_given_template_user_is_prompted():
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
