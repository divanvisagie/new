package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getDownloadUrl(userUrl string) string {
	// https://codeload.github.com/divanvisagie/postl/zip/master
	return fmt.Sprintf("https://codeload.github.com/%s/zip/master", userUrl)
}

// https://golangcode.com/unzip-files-in-go/
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			os.MkdirAll(fpath, os.ModePerm)

		} else {

			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
				return filenames, err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}

func unzipFile(filename string, projectName string) {
	files, err := Unzip(filename, projectName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unzipped: " + strings.Join(files, ", "))
}

func deleteFile(filename string) {
	os.Remove(filename)
}

func downloadFile(projectName string, url string) {

	filename := projectName + ".zip"

	out, _ := os.Create(filename)
	defer out.Close()

	response, getErr := http.Get(url)
	if getErr != nil {
		log.Fatalln(getErr.Error())
	}
	defer response.Body.Close()

	io.Copy(out, response.Body)

	unzipFile(filename, projectName)
	deleteFile(filename)
}

func main() {

	projectName := "my-project-name-i-selected"

	url := getDownloadUrl("divanvisagie/kotlin-tested-seed")

	downloadFile(projectName, url)

	fmt.Println("URL: ", url)
}
