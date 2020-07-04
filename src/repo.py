from os import system

def complete_url(url):
    if str.startswith(url, 'https:'):
        return url
    url = f'git@github.com:{url}.git'
    return url

def fetch_code(project, url):
    command = f'git clone {url}.git {project}'
    system(command)