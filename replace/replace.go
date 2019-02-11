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

// StartProcessWithString starts the procces of replacing occurences of a string
func StartProcessWithString(targetString string, targetFolder string, with string) {
	files := getAllFilePathsInDirectory(targetFolder)
	for _, filePath := range files {
		replacedIn := replaceInstancesInFile(filePath, targetString, with)
		if replacedIn != "" {
			fmt.Printf("For file %s\n", filePath)
			fmt.Printf("Replaced:\n\n%s\n\n", replacedIn)
		}
	}
}
