package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type model struct {
	textInput textinput.Model
	files     []string
	index     int
	err       errMsg
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "/path/to/music/folder"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 64

	return model{
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.files == nil {
				path := m.textInput.Value()
				files, err := getAllFilesInPath(path)
				if err != nil {
					log.Printf("Error fetching files: %v", err)
					return m, nil
				}
				m.files = files
				return m, nil
			}
			return m, m.editTags()
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.files == nil {
		return fmt.Sprintf("What’s your music folder path to edit?\n\n%s\n\n%s", m.textInput.View(), "(esc to quit)") + "\n"
	}
	return "Files loaded, press Enter to start editing tags...\n(esc to quit)\n"
}

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

func (m model) editTags() tea.Cmd {
	return func() tea.Msg {
		for _, file := range m.files {
			tag, err := id3v2.Open(file, id3v2.Options{Parse: true})
			if err != nil {
				log.Printf("Error opening file %s: %v", file, err)
				continue
			}

			fmt.Printf("Editing tags for %s\n", file)
			// Simulate tag editing process...
			tag.Close()
		}
		return errMsg(fmt.Errorf("finished editing tags"))
	}
}

func main() {
	OrganizeMusicFiles()
}

func OrganizeMusicFiles() {
	files, err := getAllFilesInPath("/home/lucas/Downloads/Rainy & Cozy Days")
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

		artistPath := filepath.Join("/home/lucas/Music", artist)
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