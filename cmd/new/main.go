package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/divanvisagie/new/internal/git"
	"github.com/divanvisagie/new/internal/prompt"
	"github.com/divanvisagie/new/internal/targ"
)

const separator = string(os.PathSeparator)

// var (
// 	app = kingpin.New("new", "generate projects from git repositories")

// 	name = app.Arg("project name", "Name of the new project").Required().String()
// 	seed = app.Arg("repository", "Custom git repo URL or GitHub <username>/<project>").Required().String()

// 	verbose = app.Flag("verbose", "Verbose mode").Short('v').Bool()
// )

func removeGitInDirectory(directoryName string) {
	gitPath := string(strings.Join([]string{directoryName, ".git"}, separator))

	dir, _ := os.Getwd()
	target := strings.Join([]string{dir, gitPath}, separator)
	err := os.RemoveAll(target)

	if err != nil {
		log.Fatalln("Failed to delete .git directory")
	}
}

func removeConfigFileInDirectory(directoryName string) {
	dir, _ := os.Getwd()
	newConfigFile := path.Join(dir, directoryName, ".new.yml")
	err := os.Remove(newConfigFile)
	if err != nil {
		log.Fatalf("Could not delete %s because %s\n", newConfigFile, err)
	}
}

func fetchRepository(seed string, name string, getUserInput func(string, string) string) {
	args := git.GetArgs(seed, name)
	err := git.RunCommand(args)
	if err != nil {
		fmt.Printf("Failed due to error: %s\n", err.Error())
	}
	removeGitInDirectory(name)
	prompt.ProcessForTarget(name, getUserInput)
	defer removeConfigFileInDirectory(name)
}

func main() {
	// prompt.Verbose = *verbose

	// fmt.Printf("verbose: %v\n", *verbose)

	// switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// default:
	// 	fetchRepository(*seed, *name, func(name string, description string) string {
	// 		reader := bufio.NewReader(os.Stdin)
	// 		fmt.Printf("\nEnter replacement text for %s\n\n    text       : %s\n    description: %s\n\n> ", name, name, description)
	// 		text, _ := reader.ReadString('\n')
	// 		text = strings.TrimSpace(text)
	// 		if text == "" {
	// 			text = name
	// 		}
	// 		return text
	// 	})
	// }

	fmt.Printf("supplied args %v\n", os.Args[1:])

	c := targ.NewContainer(os.Args[1:])
	c.Description("generate projects from git repositories")
	name := c.Arg(0).Name("project name").Description("Name of your new project").String()
	seed := c.Arg(1).Name("repository url").Description("Custom git repo URL or GitHub short: <username>/<project>").String()

	fetchRepository(seed, name, func(name string, description string) string {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\nEnter replacement text for %s\n\n    text       : %s\n    description: %s\n\n> ", name, name, description)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			text = name
		}
		return text
	})

}
