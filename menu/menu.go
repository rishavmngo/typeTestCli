package menu

import (
	"fmt"
	"strings"

	"github.com/nexidian/gocliselect"
	"golang.org/x/term"
)

func renderMenu() {

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
}
