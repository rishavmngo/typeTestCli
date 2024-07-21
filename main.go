package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	m "github.com/rishavmngo/menu-go/menu"
	"github.com/typeTest/model"
	"github.com/typeTest/ui"
	"github.com/typeTest/utils"
	"golang.org/x/term"
)

func loadWords(path string, settings model.Settings) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)

	}
	var samples []model.SampleWord
	json.Unmarshal(data, &samples)
	words := strings.Split(samples[settings.Mode].Text, samples[settings.Mode].Delimiter)
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}
	return words
}

func loadSettings(path string) model.Settings {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var settings model.Settings
	json.Unmarshal([]byte(data), &settings)
	return settings
}

var durationOfGame = 10

var timerDuration = 0

func gameStarted() bool {
	return timerDuration != 0

}
func addToInput(input *string, inp string) {
	if timerDuration == 0 {
		timerDuration = durationOfGame
		*input = ""
	}
	*input += string(inp)
}

func main() {

	// var greettingScreen bytes.Buffer
	// greettingScreen.WriteString("hello")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	configDir := filepath.Join(homeDir, ".config/typeTest-go")
	settingsPath := filepath.Join(configDir, "settings.json")
	wordsPath := filepath.Join(configDir, "words.json")

	var settings model.Settings = loadSettings(settingsPath)

	// menu.GreetingMenu(&settings, cancel)

	menu := m.NewMenu("Main Menu")

	menu.Main.Add("Play", func() {
		menu.Exit()
	})
	setting := menu.Main.Add("Settings", nil)

	menu.Main.Add("Exit", func() {
		os.Exit(0)
	})

	Mode := setting.Add("Mode", nil)
	Mode.Add("Easy", func() {
		settings.Mode = 1
		menu.Back()
	})
	Mode.Add("Advance", func() {
		settings.Mode = 2
		menu.Back()
	})
	Mode.Add("Paragraph", func() {
		settings.Mode = 0
		menu.Back()
	})

	Duration := setting.Add("Duration", nil)

	Duration.Add("10", func() {
		settings.Duration = 10
		menu.Back()
	})
	Duration.Add("30", func() {
		settings.Duration = 30
		menu.Back()
	})
	Duration.Add("60", func() {
		settings.Duration = 60
		menu.Back()
	})
	Duration.Add("120", func() {
		settings.Duration = 120
		menu.Back()
	})

	CursorCharacterMenu := setting.Add("Cursor character", nil)
	CursorCharacterMenu.Add("UnderScore Cursor(_)", func() {
		settings.CursorCharacter = "_"
		menu.Back()
	})
	CursorCharacterMenu.Add("Pipe Cursor(|)", func() {

		settings.CursorCharacter = "|"
		menu.Back()
	})

	menu.Display()

	durationOfGame = settings.Duration
	data, err := json.Marshal(settings)

	if err := os.WriteFile(settingsPath, data, 0644); err != nil {
		log.Fatalf("failed to write  file: %s", err)
	}
	var buffer bytes.Buffer

	strArray := loadWords(wordsPath, settings)

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		// term.Restore(int(os.Stdin.Fd()), oldState)
		// os.Exit(0)
		cancel()
	}()

	var input string = "Type you monkey"
	currentWord := 0

	wrongFlag := false
	// alreadyStarted := false

	inpChan := make(chan byte)

	go func() {
		defer close(inpChan)

		inpCh := make([]byte, 1)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				n, err := os.Stdin.Read(inpCh)
				if err != nil {
					fmt.Println("Error reading input:", err)
					close(inpChan)
				}
				if n > 0 {
					inpChan <- inpCh[0]
				}
			}
		}

	}()

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	cursorBlink := time.NewTicker(700 * time.Millisecond)
	defer cursorBlink.Stop()

	timerTicker := time.NewTicker(1 * time.Second)
	defer timerTicker.Stop()

	timerStr := fmt.Sprintf("%ds\r\n", durationOfGame)
mainLoop:
	for {

		if gameStarted() {
			wrongFlag = utils.CheckForTypo(input, strArray[currentWord])
		}
		select {
		case inp := <-inpChan:
			switch inp {
			case '\n':
			case '\r':
			case 3:
				cancel()
				break mainLoop
			case 127:

				inputLength := len(input)
				if inputLength > 0 && gameStarted() {
					input = input[:inputLength-1]
				}
			case 23:
				inputArray := strings.Split(input, " ")
				input = strings.Join(inputArray[:len(inputArray)-1], " ")
			case 32:
				if !wrongFlag && input == strArray[currentWord] {
					input = ""
					currentWord++
				} else {
					addToInput(&input, string(inp))
				}
			default:
				addToInput(&input, string(inp))
			}
		case <-ticker.C:
			buffer.Reset()
			ui.ClearScreen(&buffer)
			buffer.WriteString(timerStr)
			ui.RenderTextBox(&buffer, strArray, currentWord, 0, wrongFlag)
			ui.RenderInputBox(&buffer, input, cursorBlink)
			// buffer.WriteString(fmt.Sprintf("%ds", timerDuration))
			_, err := buffer.WriteTo(os.Stdout)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
			}
		case <-timerTicker.C:
			if timerDuration > 0 {
				timerDuration--
				timerStr = fmt.Sprintf("%ds\r\n", timerDuration)

				if timerDuration == 0 {
					ui.ClearScreen(&buffer)
					speed := (currentWord * 60) / durationOfGame
					fmt.Fprintf(&buffer, "Time's up!\r\nSpeed: %d WPM", speed)
					_, err := buffer.WriteTo(os.Stdout)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
					}
					// choice := menu.ExitMenu(&settings, cancel)
					choice := "exit"
					if choice == "exit" {
						ticker.Stop()
						break mainLoop
					} else {

						timerDuration = 0
						currentWord = 0
						input = ""
					}
				}
			}
		case <-ctx.Done():
			break mainLoop
		}
	}

	fmt.Println("\r\nExiting...")
}
