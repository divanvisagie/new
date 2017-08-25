package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getDownloadUrl(userUrl string) string {
	// https://codeload.github.com/divanvisagie/postl/zip/master
	return fmt.Sprintf("https://codeload.github.com/%s/zip/master", userUrl)
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
}

func main() {

	projectName := "my-project-name-i-selected"

	url := getDownloadUrl("divanvisagie/kotlin-tested-seed")

	downloadFile(projectName, url)

	fmt.Println("URL: ", url)
}
