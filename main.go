package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

const separator = string(os.PathSeparator)

func getGitArgs(projectURL string, projectName string) []string {

	dir, _ := os.Getwd()

	// https://codeload.github.com/divanvisagie/postl/zip/master
	target := strings.Join([]string{dir, projectName}, separator)
	var url string
	if strings.ContainsRune(projectURL, ':') {
		url = projectURL
	} else {
		url = fmt.Sprintf("https://github.com/%s.git", projectURL)
	}

	arguments := []string{
		"clone",
		"--depth=1",
		url,
		target,
	}

	fmt.Printf("Creating %s from %s \n", projectName, url)

	return arguments
}

func runCommand(command string, arguments []string) (string, error) {
	commandOutput, err := exec.Command(command, arguments...).Output()
	if err != nil {
		return "", err
	}
	return string(commandOutput), nil
}

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

		commandArgs := getGitArgs(*seed, *name)
		_, err := runCommand("git", commandArgs)
		if err != nil {
			fmt.Printf("Failed due to error: %s\n", err.Error())
		}

		removeGitInDirectory(*name)
	}
}
