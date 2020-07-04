import pytest

from repo import complete_url

def test_inserts_default_domain():
    actual = complete_url('divanvisagie/new')
    assert actual == 'git@github.com:divanvisagie/new.git'


