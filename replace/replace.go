package replace

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/kylelemons/godebug/diff"
)

func osFriendlyNewlineSplit(str string) []string {
	return strings.Split(strings.Replace(str, "\r\n", "\n", -1), "\n")
}

func colorDiffLine(str string) string {
	add, _ := regexp.MatchString(`^\+.*`, str)
	subtract, _ := regexp.MatchString(`^\-.*`, str)

	if add {
		return color.GreenString(str)
	}

	if subtract {
		return color.RedString(str)
	}

	return str
}

func chunkContains(chunk []string, match string) bool {
	for _, line := range chunk {
		if strings.Contains(line, match) {
			return true
		}
	}
	return false
}

func findInterestingChunks(text, match string) []string {
	var chunks []string
	lines := osFriendlyNewlineSplit(text)

	var currentChunk []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			if len(currentChunk) > 0 {
				if chunkContains(currentChunk, match) {
					chunk := strings.Join(currentChunk, "\n")
					chunks = append(chunks, chunk)
				}
				currentChunk = []string{}
			}
		} else {
			currentChunk = append(currentChunk, colorDiffLine(line))
		}
	}
	return chunks
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

func printChange(original, new, with string) {
	ans := diff.Diff(original, new)

	chunks := findInterestingChunks(ans, with)
	for _, chunk := range chunks {
		if strings.ContainsAny(chunk, "+-") {
			fmt.Printf("\n%v\n", chunk)
		}
		fmt.Printf("\n")
	}
}

func replaceInstancesInFile(targetFile string, instance string, with string) string {
	ans := ""
	bytes, err := ioutil.ReadFile(targetFile)
	if err != nil {
		fmt.Printf("Error: %s while replacing instances in file: %s", err.Error(), targetFile)
	}
	str := string(bytes)

	if strings.Contains(str, instance) {
		ans = strings.Replace(str, instance, with, -1)
	}

	if ans != "" {
		c := color.New(color.FgBlack).Add(color.BgWhite)
		output := c.Sprintf("%s :", targetFile)
		fmt.Printf("%s\n", output)
		printChange(str, ans, with)
	}

	return ans
}

func overwriteFileWith(path string, with string) {
	err := ioutil.WriteFile(path, []byte(with), 0644)
	if err != nil {
		fmt.Printf("Error %s writing to file %s", err.Error(), path)
	}
}

// StartProcessWithString starts the procces of replacing occurences of a string
func StartProcessWithString(targetString string, targetFolder string, with string) {
	files := getAllFilePathsInDirectory(targetFolder)
	for _, filePath := range files {
		replacementText := replaceInstancesInFile(filePath, targetString, with)
		if replacementText != "" {
			overwriteFileWith(filePath, replacementText)
		}
	}
}
