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

	"github.com/typeTest/ui"
	"github.com/typeTest/utils"
	"golang.org/x/term"
)

var str1 = "Paragraphs are the building blocks of papers. Many students define paragraphs in terms of length: a paragraph is a group of at least five sentences, a paragraph is half a page long, etc. In reality, though, the unity and coherence of ideas among sentences is what constitutes a paragraph"
var str2 = "hello my name is rishav and i am a computer engineer"

var durationOfGame = 10

var timerDuration = 0

func addToInput(input *string, inp string) {
	if timerDuration == 0 {
		timerDuration = durationOfGame
	}
	*input += string(inp)
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var buffer bytes.Buffer

	strArray := strings.Split(str1, " ")

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

	var input string
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
	timerTicker := time.NewTicker(1 * time.Second)
	defer timerTicker.Stop()

	timerStr := fmt.Sprintf("%ds", durationOfGame)
mainLoop:
	for {

		wrongFlag = utils.CheckForTypo(input, strArray[currentWord])
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
				if inputLength > 0 {
					input = input[:inputLength-1]
				}
			case 23:
				inputArray := strings.Split(input, " ")
				input = strings.Join(inputArray[:len(inputArray)-1], " ")
			case 32:
				if !wrongFlag {
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
			ui.RenderTextBox(&buffer, strArray, currentWord, 0, wrongFlag)
			ui.RenderInputBox(&buffer, input)
			// buffer.WriteString(fmt.Sprintf("%ds", timerDuration))
			buffer.WriteString(timerStr)
			_, err := buffer.WriteTo(os.Stdout)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
			}
		case <-timerTicker.C:
			if timerDuration > 0 {
				timerDuration--
				timerStr = fmt.Sprintf("%ds", timerDuration)

				if timerDuration == 0 {
					ticker.Stop()
					ui.ClearScreen(&buffer)
					speed := (currentWord * 60) / durationOfGame
					fmt.Fprintf(&buffer, "Time's up!\r\nSpeed: %d WPM", speed)
					_, err := buffer.WriteTo(os.Stdout)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
					}
					break mainLoop
				}
			}
		case <-ctx.Done():
			break mainLoop
		}
	}

	fmt.Println("\r\nExiting...")
}
