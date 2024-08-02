package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
)

func getAllFilesInPath(path string) ([]string, error) {
    var files []string
    err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the folder path here: ")
    text, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    text = strings.TrimSpace(text)

	files, err := getAllFilesInPath(text)
    if err != nil {
        log.Fatal(err)
    }
	album := ""
	for _, file := range files {
		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
		if err != nil {
			log.Fatal("Error while opening a file: ", err)
		}
		defer tag.Close()

		if album == "" {
			album = tag.Album()
		} else if album != tag.Album() {
			fmt.Println(album, file)
			log.Fatal("All files in the folder do not have the same album name")
		}

		log.Println(file + " | | | * * * | | | " + tag.Album())
	}
}