package git

import (
	"fmt"
	"os"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

const separator = string(os.PathSeparator)

// Args is the argument type for the git command
type Args struct {
	url    string
	target string
}

// GetArgs gets the arguments required for git
func GetArgs(url string, projectName string) *Args {
	dir, _ := os.Getwd()
	target := strings.Join([]string{dir, projectName}, separator)
	// var url string
	if !strings.ContainsRune(url, ':') {
		url = fmt.Sprintf("https://github.com/%s.git", url)
	}

	fmt.Printf("Creating %s from %s \n", projectName, url)

	return &Args{
		target: target,
		url:    url,
	}
}

// RunCommand runs the git command with the specified args
func RunCommand(args *Args) error {

	_, err := git.PlainClone(args.target, false, &git.CloneOptions{
		URL:      args.url,
		Progress: os.Stdout,
		Depth:    1,
	})

	return err
}