package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/mholt/archiver"
)

func getDownloadURL(userURL string) string {
	// https://codeload.github.com/divanvisagie/postl/zip/master
	return fmt.Sprintf("https://codeload.github.com/%s/zip/master", userURL)
}

func deleteFile(filename string) {
	filename = ".\\" + filename
	fmt.Println("Removing", filename)
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

	archiver.Zip.Open(filename, projectName)
	deleteFile(filename)
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
