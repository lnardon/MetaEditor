package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

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
    fmt.Print("Enter the folder path here: ")
    path, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    path = strings.TrimSpace(path)

	fmt.Println("Choose one of the following options:")
	fmt.Println("1. Organize music files into its respective artist and album folders")
	fmt.Println("2. Set release year")
	fmt.Println("3. Set Genre")
	fmt.Println("4. Set Artist")
	fmt.Println("5. Set Album")
	fmt.Println("6. Set Track Number")
	fmt.Println("7. Exit")
    opt, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
	fmt.Print("\nStarting...\n\n")

	switch strings.TrimSpace(opt) {
	case "1":
		OrganizeMusicFiles(path)
	case "2":
		SetReleaseYear(path)
	case "3":
		SetGenre(path)
	case "4":
		SetArtist(path)
	case "5":
		SetAlbum(path)
	case "6":
		SetTrackNumber(path)
	default:
		fmt.Println("Exiting...")
		return
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

func SetReleaseYear(path string) {
	files, err := getAllFilesInPath(path)
	if err != nil {
		log.Fatalf("Error fetching files: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the release year: ")
	year, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	year = strings.TrimSpace(year)

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
			if err != nil {
				log.Printf("Error opening file %s: %v", file, err)
				return
			}
			defer tag.Close()

			tag.SetYear(year)
			if err := tag.Save(); err != nil {
				log.Printf("Error saving file %s: %v", file, err)
				return
			}

			fmt.Printf("Set release year %s for %s\n", year, file)
			fmt.Println("^ ^ ^")
		}(file)
	}
	wg.Wait()
}

func SetGenre(path string) {
	files, err := getAllFilesInPath(path)
	if err != nil {
		log.Fatalf("Error fetching files: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the genre: ")
	genre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	genre = strings.TrimSpace(genre)

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string){
		defer wg.Done()
		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
		if err != nil {
			log.Printf("Error opening file %s: %v", file, err)
			return
		}

		tag.SetGenre(genre)
		if err := tag.Save(); err != nil {
			log.Printf("Error saving file %s: %v", file, err)
			return
		}

		tag.Close()

		fmt.Printf("Set genre %s for %s\n", genre, file)
		fmt.Println("^ ^ ^")
		}(file)
	}
	wg.Wait()
}

func SetArtist(path string) {
	files, err := getAllFilesInPath(path)
	if err != nil {
		log.Fatalf("Error fetching files: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the artist: ")
	artist, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	artist = strings.TrimSpace(artist)

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string){
		defer wg.Done()
		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
		if err != nil {
			log.Printf("Error opening file %s: %v", file, err)
			return
		}

		tag.SetArtist(artist)
		if err := tag.Save(); err != nil {
			log.Printf("Error saving file %s: %v", file, err)
			return
		}

		tag.Close()

		fmt.Printf("Set artist %s for %s\n", artist, file)
		fmt.Println("^ ^ ^")
		}(file)
	}
	wg.Wait()
}

func SetAlbum(path string) {
	files, err := getAllFilesInPath(path)
	if err != nil {
		log.Fatalf("Error fetching files: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the album: ")
	album, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	album = strings.TrimSpace(album)

	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string){
		defer wg.Done()
		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
		if err != nil {
			log.Printf("Error opening file %s: %v", file, err)
			return
		}

		tag.SetAlbum(album)
		if err := tag.Save(); err != nil {
			log.Printf("Error saving file %s: %v", file, err)
			return
		}

		tag.Close()

		fmt.Printf("Set album %s for %s\n", album, file)
		fmt.Println("^ ^ ^")
		}(file)
	}
	wg.Wait()
}

func SetTrackNumber(path string) {
	files, err := getAllFilesInPath(path)
	if err != nil {
		log.Fatalf("Error fetching files: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the total number of tracks in the album: ")
	totalAmount, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	for _, file := range files {
		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
		if err != nil {
			log.Printf("Error opening file %s: %v", file, err)
			continue
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Println(tag.Title())
		fmt.Print("Enter the track number for the track above: ")
		track, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		track = strings.TrimSpace(track)

		tag.AddTextFrame("TRCK", tag.DefaultEncoding(), fmt.Sprintf("%s/%s", track, totalAmount))
		if err := tag.Save(); err != nil {
			log.Printf("Error saving file %s: %v", file, err)
			continue
		}

		tag.Close()

		fmt.Printf("Set track number %s for %s\n", track, file)
		fmt.Println("^ ^ ^")
	}
}