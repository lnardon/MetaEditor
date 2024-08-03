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
		if !info.IsDir() && strings.HasSuffix(path, ".mp3") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func main() {
	OrganizeMusicFiles()
}

func OrganizeMusicFiles() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the folder path here: ")
    path, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    path = strings.TrimSpace(path)

	files, err := getAllFilesInPath(path)
	if err != nil {
		log.Fatalf("Error fetching files: %v", err)
	}

	for _, file := range files {
		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
		if err != nil {
			log.Printf("Error opening file %s: %v", file, err)
			continue
		}

		artist := tag.Artist()
		album := tag.Album()

		artistPath := filepath.Join(path , artist)
		if _, err := os.Stat(artistPath); os.IsNotExist(err) {
			if err := os.Mkdir(artistPath, 0755); err != nil {
				log.Printf("Error creating folder %s: %v", artistPath, err)
				continue
			}
		}

		albumPath := filepath.Join(artistPath, album)
		if _, err := os.Stat(albumPath); os.IsNotExist(err) {
			if err := os.Mkdir(albumPath, 0755); err != nil {
				log.Printf("Error creating folder %s: %v", albumPath, err)
				continue
			}
		}

		newPath := filepath.Join(albumPath, filepath.Base(file))
		if err := os.Rename(file, newPath); err != nil {
			log.Printf("Error moving file %s to %s: %v", file, newPath, err)
			continue
		}

		tag.Close()

		fmt.Printf("Moved %s to %s\n", file, newPath)
	}
}