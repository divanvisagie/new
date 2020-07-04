#!/usr/bin/env python
import argparse
import sys
from os.path import dirname, join, abspath
from repo import fetch_code

root = abspath(join(dirname(__file__), '.')) # The root of this file

    
# Command line interface

def get_options(args=sys.argv[1:]):
    parser = argparse.ArgumentParser(description='Generates new project from an existing repository')
    parser.add_argument('project', type=str, help='name of the project you want to create')
    parser.add_argument('url', type=str, help='url of the repository you want to base this project on')
    return parser.parse_args(args)

def print_repo_details(project_name, git_url):
    print(f'Project name: {project_name}\nGit URL: {git_url}')

def main():
    print('The python version of new is not yet implemented')
    options = get_options()
    print_repo_details(options.project, options.url)
    fetch_code(options.project, options.url)

if __name__ == "__main__":
    main()