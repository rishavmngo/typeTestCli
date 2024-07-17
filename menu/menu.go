package menu

import (
	"context"
	"strconv"

	"github.com/nexidian/gocliselect"
	"github.com/typeTest/model"
	"github.com/typeTest/ui"
)

func RenderMenu() string {
	menu := gocliselect.NewMenu("Main menu")
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

func GreetingMenu(settings *model.Settings, cancel context.CancelFunc) {

	ui.ClearScreenStandalone()
	choice := RenderMenu()
	switch choice {
	case "play":
	case "settings":
		ui.ClearScreenStandalone()
		choice := RenderSettingsMenu()
		if choice == "mode" {
			ui.ClearScreenStandalone()
			settings.Mode = RenderModeMenu()
			// GreetingMenu(settings, cancel)
		} else if choice == "duration" {
			ui.ClearScreenStandalone()
			settings.Duration = RenderDurationMenu()
			// GreetingMenu(settings, cancel)
		}
	case "exit":
		cancel()

	}
}

func ExitMenu(settings *model.Settings, cancel context.CancelFunc) string {

	choice := RenderMenu()
	switch choice {
	case "play":
	case "settings":
		ui.ClearScreenStandalone()
		choice := RenderSettingsMenu()
		if choice == "mode" {
			ui.ClearScreenStandalone()
			settings.Mode = RenderModeMenu()
		} else if choice == "duration" {
			ui.ClearScreenStandalone()
			settings.Duration = RenderDurationMenu()
		}
	case "exit":
		return "exit"
	}
	return "play"
}

func RenderSettingsMenu() string {

	menu := gocliselect.NewMenu("Settings")
	menu.AddItem("Mode", "mode")
	menu.AddItem("Duration", "duration")
	choice := menu.Display()
	return choice
}

func RenderModeMenu() int {

	menu := gocliselect.NewMenu("Mode")
	menu.AddItem("Paragraph", "0")
	menu.AddItem("Easy", "1")
	menu.AddItem("Advance", "2")
	choice := menu.Display()
	value, err := strconv.Atoi(choice)
	if err != nil {
		panic(err)

	}
	return value
}

func RenderDurationMenu() int {

	menu := gocliselect.NewMenu("Mode")
	menu.AddItem("2", "2")
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
