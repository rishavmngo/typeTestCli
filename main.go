package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/typeTest/menu"
	"github.com/typeTest/settings"
	s "github.com/typeTest/settings"
	"github.com/typeTest/ui"
	"github.com/typeTest/utils"
	"golang.org/x/term"
)

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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//get the setting instance
	settings := s.Get()
	//Load the settings from configuration
	settings.Load()

	//Display the greeting menu
	menu.GreetingMenu()

	//Write the seetings to configuration file
	settings.Save()

	//Set the duration and words a/q to settings
	durationOfGame = settings.Duration
	words := settings.GetWords()

	var buffer bytes.Buffer

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
	inpChan := make(chan byte)
	// controlChan := make(chan bool)
	paused := false

	go func() {
		defer close(inpChan)

		inpCh := make([]byte, 1)

		for {
			select {
			case <-ctx.Done():
				return
			default:

				if paused {
					continue
				}
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
			wrongFlag = utils.CheckForTypo(input, words[currentWord])
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
				if !wrongFlag && input == words[currentWord] {
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

			ui.MarginTop(&buffer)
			ui.MarginLeft(&buffer)
			buffer.WriteString(timerStr)
			ui.RenderTextBox(&buffer, words, currentWord, 0, wrongFlag)
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
					paused = true
					ui.ClearScreen(&buffer)
					speed := (currentWord * 60) / durationOfGame
					fmt.Fprintf(&buffer, "Time's up!\r\nSpeed: %d WPM", speed)
					_, err := buffer.WriteTo(os.Stdout)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
					}

					timerTicker.Stop()

					time.Sleep(3 * time.Second)
					fmt.Printf("\r\nEnter a key to exit!")
					inp := make([]byte, 1)
					_, _ = os.Stdin.Read(inp)

					_ = <-inpChan
					menu.ExitMenu()
					settings.Save()
					wrongFlag = false
					restart(&input, &currentWord, timerTicker, &timerStr)
					paused = false

				}
			}
		case <-ctx.Done():
			break mainLoop
		default:
		}
	}

	fmt.Println("\r\nExiting...")
}

func restart(input *string, currentWord *int, timerTicker *time.Ticker, timerStr *string) {

	*currentWord = 0
	durationOfGame = settings.Get().Duration
	timerTicker.Reset(1 * time.Second)
	timerDuration = 0
	*input = "Type you monkey"
	*timerStr = fmt.Sprintf("%ds\r\n", durationOfGame)
}
