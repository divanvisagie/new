import setuptools

with open('README.md', 'r') as fh:
    long_description = fh.read()

setuptools.setup(
    name='new-project', # Replace with your own username
    version='0.0.6',
    author='Divan Visagie',
    author_email='me@dvisagie.com',
    description='Generate new projects from git repositories',
    long_description=long_description,
    long_description_content_type='text/markdown',
    url='https://github.com/divanvisagie/new',
    packages=setuptools.find_packages(),
    classifiers=[
        'Programming Language :: Python :: 3',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
    ],
    entry_points={
        'console_scripts': ['new=new.new:main'],
    },
    python_requires='>=3.6',
)