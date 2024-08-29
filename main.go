package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	Metadata "MetaEdito/metadata"

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
	reader := bufio.NewReader(os.Stdin)
    path := os.Args[1]

	fmt.Println("Would you like to organize your music files or set metadata?(Type 'organize' or 'metadata' and press enter)")
	opt, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch strings.TrimSpace(opt) {
	case "organize":
		OrganizeMusicFiles(path)
	case "metadata":
		fmt.Println("Do you want to set metadata for all files in the folder or for a specific file?(Type 'all' or 'specific' and press enter)")
		opt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		switch strings.TrimSpace(opt) {
		case "all":
			fmt.Println("Choose one of the following options:")
			fmt.Println("1. Set release year")
			fmt.Println("2. Set Genre")
			fmt.Println("3. Set Artist")
			fmt.Println("4. Set Album")
			fmt.Println("5. Set Track Number")

			opt, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			files, err := getAllFilesInPath(path)
			if err != nil {
				log.Fatalf("Error fetching files: %v", err)
			}

			switch strings.TrimSpace(opt) {
			case "1":
				Metadata.SetReleaseYear(files)
			case "2":
				Metadata.SetGenre(files)
			case "3":
				Metadata.SetArtist(files)
			case "4":
				Metadata.SetAlbum(files)
			case "5":
				Metadata.SetTrackNumber(files)
			default:
				fmt.Println("Exiting...")
				return
			}

		case "specific":
			fmt.Println("Choose one of the following options:")
			fmt.Println("1. Set track title")
			fmt.Println("2. Set release year")
			fmt.Println("3. Set Genre")
			fmt.Println("4. Set Artist")
			fmt.Println("5. Set Album")
			fmt.Println("6. Set Track Number")

			opt, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			files, err := getAllFilesInPath(os.Args[1])
			if err != nil {
				log.Fatalf("Error fetching files: %v", err)
			}


			fmt.Println("Type the number of the file you want to set metadata for:")
			
			for i, file := range files {
				fmt.Printf("%d. %s\n", i+1, file)
			}

			fileOption, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			intOpt, err := strconv.Atoi(strings.TrimSpace(fileOption))
			if err != nil {
				fmt.Println("Error converting input to integer:", err)
				return
			}
			file := files[intOpt - 1]

			switch strings.TrimSpace(opt) {
			case "1":
				Metadata.SetTitle([]string{file})
			case "2":
				Metadata.SetReleaseYear([]string{file})
			case "3":
				Metadata.SetGenre([]string{file})
			case "4":
				Metadata.SetArtist([]string{file})
			case "5":
				Metadata.SetAlbum([]string{file})
			case "6":
				Metadata.SetTrackNumber([]string{file})
			default:
				fmt.Println("Exiting...")
				return
			}

		default:
			fmt.Println("Invalid option")
		}
	default:
		fmt.Println("Invalid option")
	}


}

func OrganizeMusicFiles(path string) {
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

		artist := strings.Split(tag.Artist(), "/")[0]
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
		fmt.Println("^ ^ ^")
	}
}