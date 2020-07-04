from os import system, chdir
import re
import shutil

def complete_url(url):
    p = re.compile('(^[A-z]*/[A-z]*)')
    if p.match(url):
        return f'git@github.com:{url}.git'
    return url

def fetch_code(project, url):
    command = f'git clone {url} {project}'
    system(command)

def ungitify(project):
    # Remove internal git from target
    shutil.rmtree(f'{project}/.git')

def reset_remotes(project):
    command = f''
    chdir(project)
    system(command)
    chdir('..')