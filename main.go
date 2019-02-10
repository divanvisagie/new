package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/divanvisagie/new/git"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const separator = string(os.PathSeparator)

func removeGitInDirectory(directoryName string) {
	path := string(strings.Join([]string{directoryName, ".git"}, separator))

	dir, _ := os.Getwd()
	target := strings.Join([]string{dir, path}, separator)
	err := os.RemoveAll(target)

	if err != nil {
		log.Fatalln("Failed to delete .git directory")
	}
}

var (
	app = kingpin.New("new", "generate projects from git repositories")

	name = app.Arg("project name", "Name of the new project").Required().String()
	seed = app.Arg("repository", "Custom git repo URL or GitHub <username>/<project>").Required().String()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	default:

		commandArgs := git.GetGitArgs(*seed, *name)
		err := git.RunCommand(commandArgs)
		if err != nil {
			fmt.Printf("Failed due to error: %s\n", err.Error())
		}

		removeGitInDirectory(*name)
	}
}
