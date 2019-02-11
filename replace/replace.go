package replace

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

// StartProcessWithString starts the procces of replacing occurences of a string
func StartProcessWithString(targetString string, targetFolder string) {

}
