# new
[![Build Status](https://img.shields.io/travis/divanvisagie/new/master.svg)](https://travis-ci.org/divanvisagie/new)
[![Release](https://img.shields.io/github/release/divanvisagie/new.svg)](https://github.com/divanvisagie/new/releases/latest)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg)](https://github.com/goreleaser)
[![Go Report Card](https://goreportcard.com/badge/github.com/divanvisagie/new)](https://goreportcard.com/report/github.com/divanvisagie/new)

Generate new projects from git repositories

new is a command to create new projects based on existing git repos. It simply shallow clones a repository to a directory with your specified new project name, and cleans up the git files like they were never there. It also supports string replacement.


## Usage

You can create new project from any github repository with just the username and project name

The following command will create a new project for you based on this one:
```sh
new myNewProject divanvisagie/new
```

You can also specify a URL to any git repository that you have access to:
```sh
new myNewProject git@gitlab.com:divan/new.git
```

These examples are just simple clones though, new also lets authors configure a 
repository so that it will prompt you to replace strings that you may want to rename
for your project.

You can give it a try with:

```sh
new myNewProject divanvisagie/kotlin-tested-seed
```

You will see that it prompts you to replace the package name, just pressing enter will skip this and keep the name:

```sh
> new myNewProject divanvisagie/kotlin-tested-seed
Creating testbed from https://github.com/divanvisagie/kotlin-tested-seed.git 
Enumerating objects: 14, done.
Counting objects: 100% (14/14), done.
Compressing objects: 100% (9/9), done.
Total 14 (delta 0), reused 10 (delta 0), pack-reused 0

Enter replacement text for com.divanvisagie.example

    text       : com.divanvisagie.example
    description: The package name

> com.divanvisagie.mynewproject
```

The reason we see this is because the `kotlin-tested-seed` repository contains a `.new.yml` file at its root:

```yml
replace:
  strings:
    - match: com.divanvisagie.example
      description: The package name 
```

You can configure as many match strings as you want in your own seed projects.

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

### Development

```sh
goreleaser --skip-validate --skip-publish --rm-dist
```
