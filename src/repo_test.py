import pytest

from repo import complete_url

def test_inserts_default_domain_given_github_short():
    actual = complete_url('divanvisagie/new')
    assert actual == 'git@github.com:divanvisagie/new.git'

def test_does_not_insert_when_given_long():
    actual = complete_url('git@github.com:divanvisagie/new.git')
    assert actual == 'git@github.com:divanvisagie/new.git'

