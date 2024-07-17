package menu

import (
	"strconv"

	"github.com/nexidian/gocliselect"
)

func RenderMenu() string {
	menu := gocliselect.NewMenu("Choose a option")
	menu.AddItem("Play", "play")
	menu.AddItem("Settings", "settings")
	menu.AddItem("Exit", "exit")
	choice := menu.Display()
	return choice

	// width, height, _ := term.GetSize(0)
	// paddingX := (width - len("center")) / 2
	// paddingY := (height) / 2
	// fmt.Print(strings.Repeat("\n", paddingY))
	// fmt.Print(strings.Repeat(" ", paddingX))
	// fmt.Print("Center\n")
}

func RenderSettingsMenu() string {

	menu := gocliselect.NewMenu("Settings")
	menu.AddItem("Mode", "mode")
	menu.AddItem("Duration", "duration")
	choice := menu.Display()
	return choice
}

func RenderModeMenu() string {

	menu := gocliselect.NewMenu("Mode")
	menu.AddItem("Easy", "easy")
	menu.AddItem("Advance", "advance")
	menu.AddItem("Paragraph", "paragraph")
	choice := menu.Display()
	return choice
}

func RenderDurationMenu() int {

	menu := gocliselect.NewMenu("Mode")
	menu.AddItem("15", "15")
	menu.AddItem("20", "20")
	menu.AddItem("30", "30")
	menu.AddItem("60", "60")
	menu.AddItem("120", "120")
	menu.AddItem("300", "300")
	choice := menu.Display()
	value, err := strconv.Atoi(choice)
	if err != nil {
		panic(err)

	}
	return value
}
