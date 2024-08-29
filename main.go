package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	Metadata "MetaEditor/metadata"
	Utils "MetaEditor/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
    path := os.Args[1]

	fmt.Println("Ћ - Would you like to organize your music files or set metadata?")
	fmt.Println("(Type 'organize' or 'metadata' and press enter)")

	module, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch strings.TrimSpace(module) {
	case "organize":
		files, err := Utils.GetAllFilesInPath(path)
		if err != nil {
			log.Fatalf("Error fetching files: %v", err)
		}

		Utils.OrganizeMusicFiles(files, path)

	case "metadata":
		fmt.Println("\nЋ - Do you want to set metadata for all files in the folder or for a specific file?")
		fmt.Println("(Type 'all' or 'specific' and press enter)")
		
		target, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		switch strings.TrimSpace(target) {
		case "all":
			fmt.Println("\nЋ - Choose one of the following options:")
			fmt.Println(" ☞ 1. Set release year")
			fmt.Println(" ☞ 2. Set Genre")
			fmt.Println(" ☞ 3. Set Artist")
			fmt.Println(" ☞ 4. Set Album")
			fmt.Println(" ☞ 5. Set Track Number")

			opt, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			files, err := Utils.GetAllFilesInPath(path)
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
			fmt.Println("\nЋ - Choose one of the following options:")
			fmt.Println(" ☞ 1. Set track title")
			fmt.Println(" ☞ 2. Set release year")
			fmt.Println(" ☞ 3. Set Genre")
			fmt.Println(" ☞ 4. Set Artist")
			fmt.Println(" ☞ 5. Set Album")
			fmt.Println(" ☞ 6. Set Track Number")

			opt, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			files, err := Utils.GetAllFilesInPath(os.Args[1])
			if err != nil {
				log.Fatalf("Error fetching files: %v", err)
			}


			fmt.Println("\nЋ - Type the number of the file you want to set metadata for:")
			
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
