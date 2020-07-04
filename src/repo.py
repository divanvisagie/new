def complete_url(url):
    if str.startswith(url, 'https:'):
        return url
    url = f'git@github.com:{url}.git'
    return url
