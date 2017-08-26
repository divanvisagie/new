package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Operating system path seperator;
const SEPERATOR = string(os.PathSeparator)

func getDownloadURL(userURL string) string {
	// https://codeload.github.com/divanvisagie/postl/zip/master
	return fmt.Sprintf("https://codeload.github.com/%s/tar.gz/master", userURL)
}

func deleteFile(filename string) {
	filename = ".\\" + filename
	fmt.Println("Removing", filename)
	os.Remove(filename)
}

func containsMaster(name string) {
}

func cleanPath(path string) string {
	split := strings.Split(path, SEPERATOR)
	if strings.Contains(split[1], "master") {
		fmt.Println(">>>>")
		split = append(split[:1], split[2:]...)
	}

	return strings.Join(split, SEPERATOR)
}

// Untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func Untar(dst string, r io.Reader) error {

	gzr, err := gzip.NewReader(r)
	defer gzr.Close()
	if err != nil {
		return err
	}

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		fmt.Println("dst:", dst)

		// the target location where the dir/file should be created
		target := filepath.Join(dst, header.Name)

		target = cleanPath(target)

		fmt.Println("target: ", target)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer f.Close()

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}
		}
	}
}

func downloadFile(projectName string, url string) {

	response, getErr := http.Get(url)
	if getErr != nil {
		log.Fatalln(getErr.Error())
	}
	defer response.Body.Close()

	Untar(projectName, response.Body)

}

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("You need to pass in parameters")
	}

	if len(args) == 1 {
		log.Fatalln("you need to provide a seed url")
	}

	projectName := args[0]
	url := getDownloadURL(args[1])

	downloadFile(projectName, url)

	fmt.Println("URL: ", url)
}
