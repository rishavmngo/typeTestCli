package main

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/nexidian/gocliselect"
	"github.com/typeTest/ui"
	"golang.org/x/term"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func main() {

	var buffer bytes.Buffer

	strArray := strings.Split("Paragraphs are the building blocks of papers. Many students define paragraphs in terms of length: a paragraph is a group of at least five sentences, a paragraph is half a page long, etc. In reality, though, the unity and coherence of ideas among sentences is what constitutes a paragraph", " ")
	ui.RenderTextBox(&buffer, strArray, 35, 0)
	_, err := buffer.WriteTo(os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing buffer to stdout: %v\n", err)
	}

	menu := gocliselect.NewMenu("Chose a colour")
	menu.AddItem("Red", "red")
	menu.AddItem("Blue", "blue")
	menu.AddItem("Green", "green")
	menu.AddItem("Yellow", "yellow")
	menu.AddItem("Cyan", "cyan")
	width, height, _ := term.GetSize(0)
	paddingX := (width - len("center")) / 2
	paddingY := (height) / 2
	fmt.Print(strings.Repeat("\n", paddingY))
	fmt.Print(strings.Repeat(" ", paddingX))
	fmt.Print("Center\n")

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

	return
	fmt.Println("Terminal is now in raw mode. Type 'exit' to quit.")
	buf := make([]byte, 1)
	var input string

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
		if n > 0 {
			char := buf[0]
			if char == '\r' || char == '\n' {
				// fmt.Println()
				fmt.Printf("%s", "\n\r")
				if input == "exit" {
					break
				}
				input = ""
			} else {
				input += string(char)
				fmt.Print(string(char))
			}
		}
	}

	fmt.Println("Exiting...")
}
