package replace

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getAllFilePathsInDirectory(targetFolder string) []string {
	var paths []string
	err := filepath.Walk(targetFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			if !info.IsDir() {
				paths = append(paths, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return paths
}

func replaceInstancesInFile(targetFile string, instance string, with string) string {
	ans := ""
	bytes, err := ioutil.ReadFile(targetFile)
	if err != nil {
		fmt.Printf("Error: %s while replacing instances in file: %s", err.Error(), targetFile)
	}
	str := string(bytes)

	fmt.Println(str)

	if strings.Contains(str, instance) {
		ans = strings.Replace(str, instance, with, -1)
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
			fmt.Printf("For file %s replacing '%s' with '%s'\n", filePath, targetString, with)
			fmt.Printf("Replaced:\n\n%s\n\n", replacementText)

			overwriteFileWith(filePath, replacementText)
		}
	}
}
