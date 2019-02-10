# new
Generate new projects from git repositories

new is a command to create new projects based on existing git repos. It simply shallow clones a repository to a directory with your specified new project name, and cleans up the git files like they were never there.

Future plans include support for name replacement and anything else that seems useful, feedback is appreciated.

[![Build Status](https://travis-ci.org/divanvisagie/new.svg?branch=master)](https://travis-ci.org/divanvisagie/new)


## Installation Options

### Go get
If you have golang installed , you can simply run the following on any platform:

```sh
go get github.com/divanvisagie/new
```

### Windows 

First install [scoop](http://scoop.sh/)

```sh
scoop bucket add divanvisagie https://github.com/divanvisagie/scoop-bucket
scoop install new
```

### macOS

First install [homebrew](https://brew.sh/)

```sh
brew tap divanvisagie/homebrew-tap
brew install divanvisagie/homebrew-tap/new
```

### Linux

Download the appropriate package from [here](https://github.com/divanvisagie/new/releases).

Use either the debian package or the tarball

### Manual Installation

Download the latest [tar.gz](https://github.com/divanvisagie/new/releases) and run `install.sh`

## Run

To run , simply type the new command, the name of the project you want to generate and the github repo name that you want to seed from

```sh
new myProjectName divanvisagie/kotlin-tested-seed
```
