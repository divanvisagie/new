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

// GetGitArgs gets the arguments required for git
func GetGitArgs(projectURL string, projectName string) *Args {

	dir, _ := os.Getwd()

	target := strings.Join([]string{dir, projectName}, separator)
	var url string
	if strings.ContainsRune(projectURL, ':') {
		url = projectURL
	} else {
		url = fmt.Sprintf("https://github.com/%s.git", projectURL)
	}

	// arguments := []string{
	// 	"clone",
	// 	"--depth=1",
	// 	url,
	// 	target,
	// }

	fmt.Printf("Creating %s from %s \n", projectName, url)

	return &Args{
		target: target,
		url:    url,
	}
}

func RunCommand(args *Args) error {

	_, err := git.PlainClone(args.target, false, &git.CloneOptions{
		URL:      args.url,
		Progress: os.Stdout,
		Depth:    1,
	})

	return err
}
