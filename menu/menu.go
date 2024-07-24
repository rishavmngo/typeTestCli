package menu

import (
	"fmt"
	"os"

	m "github.com/rishavmngo/menu-go/menu"
	s "github.com/typeTest/settings"
)

func GreetingMenu() {

	settings := s.Get()

	menu := m.NewMenu("Main Menu")

	menu.Main.Add("Play", func() {
		menu.Exit()
	})
	settingsMenu := menu.Main.Add("Settings", nil)

	menu.Main.Add("Exit", func() {
		os.Exit(0)
	})

	Mode := settingsMenu.Add("Mode", nil)

	for _, mode := range settings.GetModeList() {
		Mode.Add(mode, func() {
			settings.Mode = mode
			menu.Back()
		})

	}

	Duration := settingsMenu.Add("Duration", nil)

	// Duration.Add("2", func() {
	// 	settings.Duration = 2
	// 	menu.Back()
	// })
	Duration.Add("10", func() {
		settings.Duration = 10
		menu.Back()
	})

	Duration.Add("15", func() {
		settings.Duration = 15
		menu.Back()
	})
	Duration.Add("30", func() {
		settings.Duration = 30
		menu.Back()
	})
	Duration.Add("60", func() {
		settings.Duration = 60
		menu.Back()
	})
	Duration.Add("120", func() {
		settings.Duration = 120
		menu.Back()
	})

	CursorCharacterMenu := settingsMenu.Add("Cursor character", nil)
	CursorCharacterMenu.Add("UnderScore Cursor(_)", func() {
		settings.CursorCharacter = "_"
		menu.Back()
	})
	CursorCharacterMenu.Add("Pipe Cursor(|)", func() {

		settings.CursorCharacter = "█"
		menu.Back()
	})

	CursorCharacterMenu.Add(fmt.Sprintf("Custom (%s)", settings.CursorCharacter), func() {

		inp := make([]byte, 1)

		fmt.Println("Paste your character:")
		n, err := os.Stdin.Read(inp)
		if err != nil {
			panic(err)
		}

		if n > 0 {
			settings.CursorCharacter = string(inp[0])
		}
		menu.Back()
	})
	menu.Display()
}

func ExitMenu() {

	settings := s.Get()

	menu := m.NewMenu("Main Menu")

	menu.Main.Add("Play again", func() {
		menu.Exit()
	})
	settingsMenu := menu.Main.Add("Settings", nil)

	menu.Main.Add("Exit", func() {
		os.Exit(0)
	})

	Mode := settingsMenu.Add("Mode", nil)

	for _, mode := range settings.GetModeList() {
		Mode.Add(mode, func() {
			settings.Mode = mode
			menu.Back()
		})

	}
	Duration := settingsMenu.Add("Duration", nil)

	Duration.Add("10", func() {
		settings.Duration = 10
		menu.Back()
	})

	Duration.Add("15", func() {
		settings.Duration = 15
		menu.Back()
	})
	Duration.Add("30", func() {
		settings.Duration = 30
		menu.Back()
	})
	Duration.Add("60", func() {
		settings.Duration = 60
		menu.Back()
	})
	Duration.Add("120", func() {
		settings.Duration = 120
		menu.Back()
	})

	CursorCharacterMenu := settingsMenu.Add("Cursor character", nil)
	CursorCharacterMenu.Add("UnderScore Cursor(_)", func() {
		settings.CursorCharacter = "_"
		menu.Back()
	})
	CursorCharacterMenu.Add("Pipe Cursor(|)", func() {

		settings.CursorCharacter = "█"
		menu.Back()
	})

	menu.Display()
}
