package ui

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/typeTest/settings"
	"golang.org/x/term"
)

var leftTopCorner = "╔"
var leftBottomCorner = "╚"
var rightTopCorner = "╗"
var rightBottomCorner = "╝"
var horizontalLine = "═"
var verticalLine = "║"
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Black = "\033[30m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var GrayBackground = "\033[47m"
var White = "\033[97m"
var reset = "\033[0m"

var width = 60
var termWidth, height, _ = term.GetSize(0)
var paddingX = (termWidth - width) / 2
var paddingY = (height) / 8

func getMiddleOfTheArray(startIndex, endIndex int) int {

	return startIndex + (endIndex-startIndex)/2

}

func printCurrentWord(buffer *bytes.Buffer, word string) {
	buffer.WriteString(fmt.Sprintf(GrayBackground + Black + word + reset))

}
func printAttemptedWord(buffer *bytes.Buffer, word string) {

	buffer.WriteString(fmt.Sprintf(Green + word + reset))

}

func printWrongWord(buffer *bytes.Buffer, word string) {
	buffer.WriteString(fmt.Sprintf(Red + word + reset))
}

func getChunkRange(words []string, currentIndex, W int) (int, int) {

	T := len(words)
	chunkNumber := int(math.Floor(float64(currentIndex) / float64(W)))

	startIndex := chunkNumber * W
	endIndex := int(math.Min(float64(startIndex+W-1), float64(T-1)))
	return startIndex, endIndex
}

func getLineBreak(startIndex, endIndex int) int {

	return getMiddleOfTheArray(startIndex, endIndex)
}
func printSpace(buffer *bytes.Buffer, noOfSpaces int) {

	for i := 0; i < noOfSpaces; i++ {
		buffer.WriteString(" ")
	}
}

func printEnclosedBox(buffer *bytes.Buffer, text []string, currentWord, width int, wrongFlag bool) {

	printTopBox(buffer, width)

	lineWidth := width
	MarginLeft(buffer)
	buffer.WriteString(verticalLine)
	for pointer := 0; pointer < len(text); {
		word := text[pointer]
		if lineWidth-(len(word)+1) >= 0 {
			printWord(buffer, text, pointer, currentWord, wrongFlag)
			buffer.WriteString(" ")
			lineWidth -= (len(word) + 1)
			pointer++
		} else {

			printSpace(buffer, lineWidth)
			lineWidth = width
			buffer.WriteString(fmt.Sprintf("%s\r\n", verticalLine))
			MarginLeft(buffer)
			buffer.WriteString(fmt.Sprintf("%s", verticalLine))

		}
	}
	printSpace(buffer, lineWidth)
	buffer.WriteString(verticalLine)
	buffer.WriteString("\r\n")

	printBottomBox(buffer, width)
}

func printWord(buffer *bytes.Buffer, text []string, pointer, currentWord int, wrongFlag bool) {
	if pointer == currentWord {
		if wrongFlag {
			printWrongWord(buffer, text[pointer])

		} else {

			printCurrentWord(buffer, text[pointer])
		}
	} else if pointer < currentWord {
		printAttemptedWord(buffer, text[pointer])
	} else {
		buffer.WriteString(text[pointer])
	}
}
func printTopBox(buffer *bytes.Buffer, width int) {

	MarginLeft(buffer)
	buffer.WriteString(leftTopCorner)
	buffer.WriteString(strings.Repeat(horizontalLine, width))
	buffer.WriteString(rightTopCorner)
	buffer.WriteString("\r\n")
}
func printBottomBox(buffer *bytes.Buffer, width int) {

	MarginLeft(buffer)
	buffer.WriteString(leftBottomCorner)
	buffer.WriteString(strings.Repeat(horizontalLine, width))
	buffer.WriteString(rightBottomCorner)
	buffer.WriteString("\r\n")
}

func RenderTextBox(buffer *bytes.Buffer, text []string, currentWord, currentLetter int, wrongFlag bool) {
	windowSize := 25

	startIndex, endIndex := getChunkRange(text, currentWord, windowSize)

	printEnclosedBox(buffer, text[startIndex:endIndex+1], (currentWord - startIndex), width, wrongFlag)
	buffer.WriteString("\r\n")

}

var cursor = settings.Get().CursorCharacter

var hide bool = true

func RenderInputBox(buffer *bytes.Buffer, text string, cursorBlink *time.Ticker) {

	printTopBox(buffer, width)

	MarginLeft(buffer)
	// buffer.WriteString(strings.Repeat(" ", paddingX))
	buffer.WriteString(verticalLine)
	buffer.WriteString(text)
	select {
	case <-cursorBlink.C:
		if hide {
			cursor = " "
			hide = false

		} else {
			cursor = settings.Get().CursorCharacter
			hide = true
		}
	default:
	}
	buffer.WriteString(cursor)
	buffer.WriteString(strings.Repeat(" ", width-(len(text)+1)))
	buffer.WriteString(verticalLine)
	buffer.WriteString("\r\n")
	printBottomBox(buffer, width)

}

func ClearScreen(buffer *bytes.Buffer) {
	buffer.WriteString("\033[2J") // Clear the screen
	buffer.WriteString("\033[H")  // Move the cursor to the top-left corner
}

func ClearScreenStandalone() {

	fmt.Print("\033[2J") // Clear the screen
	fmt.Print("\033[H")  // Move the cursor to the top-left corner
}

func MarginTop(buffer *bytes.Buffer) {

	buffer.WriteString(strings.Repeat("\n", paddingY))
}

func MarginLeft(buffer *bytes.Buffer) {

	buffer.WriteString(strings.Repeat(" ", paddingX))
}
