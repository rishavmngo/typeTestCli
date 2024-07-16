package ui

import (
	"bytes"
	"fmt"
	"math"
	"strings"
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

func getMiddleOfTheArray(startIndex, endIndex int) int {

	return startIndex + (endIndex-startIndex)/2

}

func printCurrentWord(buffer *bytes.Buffer, word string) {
	fmt.Print(GrayBackground + Black + word + reset)
	buffer.WriteString(fmt.Sprintf(GrayBackground + Black + word + reset))

}
func printAttemptedWord(buffer *bytes.Buffer, word string) {

	fmt.Print(Green + word + reset)
	buffer.WriteString(fmt.Sprintf(Green + word + reset))

}

func printWrongWord(buffer *bytes.Buffer, word string) {
	fmt.Print(Red + word + reset)
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
		fmt.Print(" ")
		buffer.WriteString(" ")
	}
}

func printEnclosedBox(buffer *bytes.Buffer, text []string, currentWord, width int) {

	printTopBox(buffer, width)

	lineWidth := width
	fmt.Print(verticalLine)
	buffer.WriteString(verticalLine)
	for pointer := 0; pointer < len(text); {
		word := text[pointer]
		if lineWidth-(len(word)+1) >= 0 {
			printWord(buffer, text, pointer, currentWord)
			fmt.Print(" ")
			buffer.WriteString(" ")
			lineWidth -= (len(word) + 1)
			pointer++
		} else {

			printSpace(buffer, lineWidth)
			lineWidth = width
			fmt.Printf("%s\n%s", verticalLine, verticalLine)
			buffer.WriteString(fmt.Sprintf("%s\n%s", verticalLine, verticalLine))
		}
	}
	printSpace(buffer, lineWidth)
	fmt.Print(verticalLine)
	fmt.Print("\n")
	buffer.WriteString(verticalLine)
	buffer.WriteString("\n")

	printBottomBox(buffer, width)
}

func printWord(buffer *bytes.Buffer, text []string, pointer, currentWord int) {
	if pointer == currentWord {
		printCurrentWord(buffer, text[pointer])
	} else if pointer < currentWord {
		printAttemptedWord(buffer, text[pointer])
	} else {
		fmt.Print(text[pointer])
		buffer.WriteString(text[pointer])
	}
}
func printTopBox(buffer *bytes.Buffer, width int) {
	fmt.Print(leftTopCorner)
	fmt.Print(strings.Repeat(horizontalLine, width))
	fmt.Print(rightTopCorner)
	fmt.Print("\n")
	buffer.WriteString(leftTopCorner)
	buffer.WriteString(strings.Repeat(horizontalLine, width))
	buffer.WriteString(rightTopCorner)
	buffer.WriteString("\n")
}
func printBottomBox(buffer *bytes.Buffer, width int) {

	fmt.Print(leftBottomCorner)
	fmt.Print(strings.Repeat(horizontalLine, width))
	fmt.Print(rightBottomCorner)
	fmt.Print("\n")
	buffer.WriteString(leftBottomCorner)
	buffer.WriteString(strings.Repeat(horizontalLine, width))
	buffer.WriteString(rightBottomCorner)
	buffer.WriteString("\n")
}

func RenderTextBox(buffer *bytes.Buffer, text []string, currentWord, currentLetter int) {
	windowSize := 25

	startIndex, endIndex := getChunkRange(text, currentWord, windowSize)

	width := 60
	printEnclosedBox(buffer, text[startIndex:endIndex+1], (currentWord - startIndex), width)
	fmt.Print("\n")
	buffer.WriteString("\n")

}
