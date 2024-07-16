package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/typeTest/menu"
	"github.com/typeTest/ui"
	"golang.org/x/term"
)

func clearScreen(buffer *bytes.Buffer) {
	buffer.WriteString("\033[2J") // Clear the screen
	buffer.WriteString("\033[H")  // Move the cursor to the top-left corner
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func checkForEnd(currentWord, lenOfText int) bool {

	return currentWord == lenOfText

}

func checkForTypo(input, str string) bool {

	if len(input) <= len(str) && input == str[:len(input)] {
		return false
	} else {
		return true
	}

}

var str1 = "Paragraphs are the building blocks of papers. Many students define paragraphs in terms of length: a paragraph is a group of at least five sentences, a paragraph is half a page long, etc. In reality, though, the unity and coherence of ideas among sentences is what constitutes a paragraph"
var str2 = "hello my name is rishav and i am a computer engineer"

func main() {

	var startTime time.Time

	var buffer bytes.Buffer

	strArray := strings.Split(str2, " ")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()

	inpChar := make([]byte, 1)
	var input string

	currentWord := 0

	wrongFlag := false

	alreadyStarted := false
	for {
		// buffer.Reset()
		clearScreen(&buffer)
		ui.RenderTextBox(&buffer, strArray, currentWord, 0, wrongFlag)
		ui.RenderInputBox(&buffer, input)
		_, err := buffer.WriteTo(os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
		}
		n, err := os.Stdin.Read(inpChar)
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		if n > 0 {
			if !alreadyStarted {

				startTime = time.Now()
				alreadyStarted = true
			}
			if inpChar[0] == 3 {
				break
			}
			if input == "exit" {
				break
			}
			if inpChar[0] == '\r' || inpChar[0] == '\n' {
			} else if inpChar[0] == 127 {

				inputLength := len(input)
				if inputLength > 0 {
					input = input[:inputLength-1]
				}

			} else if inpChar[0] == 23 {

				inputArray := strings.Split(input, " ")
				input = strings.Join(inputArray[:len(inputArray)-1], " ")

			} else if inpChar[0] == 32 {
				if !wrongFlag {
					input = ""
					currentWord++
				} else {

					input += string(inpChar[0])
				}
			} else {
				input += string(inpChar[0])
			}

			if checkForEnd(currentWord, len(strArray)) {
				duration := time.Since(startTime)

				clearScreen(&buffer)
				buffer.WriteString(fmt.Sprintf("Your speed is %.2f WPM\r\n", math.Round(float64(currentWord-1)/duration.Minutes())))

				_, err := buffer.WriteTo(os.Stdout)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
				}
				choice := menu.RenderMenu()
				if choice == "restart" {
					input = ""
					currentWord = 0
					alreadyStarted = false
					wrongFlag = false
					clearScreen(&buffer)
				} else if choice == "exit" {
					break
				}
			}
			wrongFlag = checkForTypo(input, strArray[currentWord])
		}
	}

	fmt.Println("Exiting...")
}
