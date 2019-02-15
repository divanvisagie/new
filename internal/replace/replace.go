package replace

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/divanvisagie/new/internal/printer"
	"github.com/fatih/color"
)

// Verbose determines if we should print more stuff
var Verbose = false

// Replacement represents a string that needs to be replaced and it's metadata
type Replacement struct {
	Match string
	With  string
}

func chunkContains(chunk []string, match string) bool {
	for _, line := range chunk {
		if strings.Contains(line, match) {
			return true
		}
	}
	return false
}

func shouldIgnore(info os.FileInfo) bool {
	return info.IsDir() || info.Name() == ".new.yml"
}

func getAllFilePathsInDirectory(targetFolder string) []string {
	var paths []string
	err := filepath.Walk(targetFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !shouldIgnore(info) {
				paths = append(paths, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return paths
}

func replaceInstancesInFile(targetFile string, replacements *[]Replacement) string {
	newText := ""
	bytes, err := ioutil.ReadFile(targetFile)
	if err != nil {
		fmt.Printf("Error: %s while replacing instances in file: %s", err.Error(), targetFile)
	}
	originalText := string(bytes)
	workingText := originalText

	for _, r := range *replacements {
		if strings.Contains(workingText, r.Match) {
			newText = strings.Replace(workingText, r.Match, r.With, -1)
			workingText = newText
		}
	}
	if Verbose {
		if newText != "" {
			c := color.New(color.FgBlack).Add(color.BgWhite)
			output := c.Sprintf("%s :", targetFile)
			fmt.Printf("%s\n", output)

			printer.PrintChange(originalText, newText)
		}
	}

	return newText
}

func overwriteFileWith(path string, with string) {
	err := ioutil.WriteFile(path, []byte(with), 0644)
	if err != nil {
		fmt.Printf("Error %s writing to file %s", err.Error(), path)
	}
}

// StartReplacementProcess starts the process of replacing collected strings
func StartReplacementProcess(replacements *[]Replacement, targetDirectory string) {
	files := getAllFilePathsInDirectory(targetDirectory)

	for _, filePath := range files {
		replacementText := replaceInstancesInFile(filePath, replacements)
		if replacementText != "" {
			overwriteFileWith(filePath, replacementText)
		}
	}
}
