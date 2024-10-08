package metadata

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/bogem/id3v2/v2"
)

func SetTitle(files []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Ћ - Enter the title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	title = strings.TrimSpace(title)
	fmt.Println("")

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

			tag.SetTitle(title)
			if err := tag.Save(); err != nil {
				log.Printf("Error saving file %s: %v", file, err)
				return
			}

			fmt.Printf("✅ Set title %s for %s\n", title, file)
		}(file)
	}

	wg.Wait()
}

func SetReleaseYear(files []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Ћ - Enter the release year: ")
	year, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	year = strings.TrimSpace(year)
	fmt.Println("")

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

			fmt.Printf("✅ Set release year %s for %s\n", year, file)
		}(file)
	}
	wg.Wait()
}

func SetGenre(files []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Ћ - Enter the genre: ")
	genre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	genre = strings.TrimSpace(genre)
	fmt.Println("")

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

			fmt.Printf("✅ Set genre %s for %s\n", genre, file)
		}(file)
	}
	wg.Wait()
}

func SetArtist(files []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Ћ - Enter the artist: ")
	artist, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	artist = strings.TrimSpace(artist)
	fmt.Println("")

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

			fmt.Printf("✅ Set artist %s for %s\n", artist, file)
		}(file)
	}
	wg.Wait()
}

func SetAlbum(files []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Ћ - Enter the album: ")
	album, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	album = strings.TrimSpace(album)
	fmt.Println("")

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

			fmt.Printf("✅ Set album %s for %s\n", album, file)
		}(file)
	}
	wg.Wait()
}

func SetTrackNumber(files []string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Ћ - What is the total number of tracks in the album: ")
	totalAmount, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("")

	for _, file := range files {
		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
		if err != nil {
			log.Printf("Error opening file %s: %v", file, err)
			continue
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\n Ћ - Enter the track number for the track below: \n")
		fmt.Println(tag.Title())
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

		fmt.Println("")
		fmt.Printf("✅ Set track number %s for %s\n", track, file)
	}
}


// func GetCover(path string) {
// 	files, err := getAllFilesInPath(path)
// 	if err != nil {
// 		log.Fatalf("Error fetching files: %v", err)
// 	}

// 	for _, file := range files {
// 		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
// 		if err != nil {
// 			log.Printf("Error opening file %s: %v", file, err)
// 			continue
// 		}

// 		frames := tag.GetFrames(tag.CommonID("Attached picture"))
// 		for _, f := range frames {
// 			pic, ok := f.(id3v2.PictureFrame)
// 			if ok {
// 				filename := fmt.Sprintf("cover_%s.jpg", tag.Album())
// 				if err := os.WriteFile(filename, pic.Picture, 0644); err != nil {
// 					log.Printf("Error saving file %s: %v", filename, err)
// 					continue
// 				}
// 				log.Printf("Cover image saved as %s", filename)
// 			}
// 		}
// 	}
// }

// func AttachCover(path string) {
// 	files, err := getAllFilesInPath(path)
// 	if err != nil {
// 		log.Fatalf("Error fetching files: %v", err)
// 	}

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter the path to the cover image: ")
// 	coverPath, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}
// 	coverPath = strings.TrimSpace(coverPath)

// 	for _, file := range files {
// 		tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
// 		if err != nil {
// 			log.Printf("Error opening file %s: %v", file, err)
// 			continue
// 		}

// 		cover, err := os.ReadFile(coverPath)
// 		if err != nil {
// 			log.Printf("Error reading file %s: %v", coverPath, err)
// 			continue
// 		}

// 		tag.AddAttachedPicture(id3v2.PictureFrame{
// 			Encoding:    id3v2.EncodingUTF8,
// 			MimeType:    "image/jpeg",
// 			PictureType: id3v2.PTFrontCover,
// 			Description: "Front cover",
// 			Picture:     cover,
// 		})
// 		if err := tag.Save(); err != nil {
// 			log.Printf("Error saving file %s: %v", file, err)
// 			continue
// 		}

// 		tag.Close()

// 		fmt.Printf("Saved cover image to %s", file)
// 	}
// }