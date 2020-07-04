#!/usr/bin/env python
import argparse
import sys
from os.path import dirname, join, abspath
from new.repo import fetch_code, ungitify, complete_url
from new.template import read

root = abspath(join(dirname(__file__), '.')) # The root of this file

def get_options(args=sys.argv[1:]):
    parser = argparse.ArgumentParser(description='Generates new project from an existing repository')
    parser.add_argument('project', type=str, help='name of the project you want to create')
    parser.add_argument('url', type=str, help='url of the repository you want to base this project on')
    parser.add_argument('-p', '--preserve', dest='preserve', action='store_true' ,help='preserve the git history of the template')
    return parser.parse_args(args)

def print_repo_details(project_name, git_url):
    print(f'Project name: {project_name}\nGit URL: {git_url}')

def main():
    options = get_options()
    url = complete_url(options.url)

    print_repo_details(options.project, url)
    fetch_code(options.project, url)
    if not options.preserve:
        ungitify(options.project)
    else:
        print('Preserving git history')
    read(options.project)
    

if __name__ == "__main__":
    main()
