
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
	"github.com/typeTest/utils"
	"golang.org/x/term"
)

var str1 = "Paragraphs are the building blocks of papers. Many students define paragraphs in terms of length: a paragraph is a group of at least five sentences, a paragraph is half a page long, etc. In reality, though, the unity and coherence of ideas among sentences is what constitutes a paragraph"
var str2 = "hello my name is rishav and i am a computer engineer"

const durationInSeconds = 40

func main() {
	timerDuration := 60

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

	inpChan := make(chan byte)

	go func() {

		inpCh := make([]byte, 1)

		for {
			n, err := os.Stdin.Read(inpCh)
			if err != nil {
				fmt.Println("Error reading input:", err)
				close(inpChan)
			}

			if n > 0 {
				inpChan <- inpCh[0]

			}

		}

	}()

	for {
		ui.ClearScreen(&buffer)
		ui.RenderTextBox(&buffer, strArray, currentWord, 0, wrongFlag)
		ui.RenderInputBox(&buffer, input)
		buffer.WriteString(fmt.Sprintf("%ds", timerDuration))

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

			if utils.CheckForEnd(currentWord, len(strArray)) {
				duration := time.Since(startTime)

				ui.ClearScreen(&buffer)
				buffer.WriteString(fmt.Sprintf("Your speed is %.2f WPM\r\n", math.Round(float64(currentWord-1)/duration.Minutes())))

				_, err := buffer.WriteTo(os.Stdout)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
				}
				choice := menu.RenderMenu()
				if choice == "play" {
					input = ""
					currentWord = 0
					alreadyStarted = false
					wrongFlag = false
					ui.ClearScreen(&buffer)
				} else if choice == "exit" {
					break
				}
			}
			wrongFlag = utils.CheckForTypo(input, strArray[currentWord])
		}
	}

	fmt.Println("Exiting...")
}
