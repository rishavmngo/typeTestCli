package menu

import (
	"github.com/nexidian/gocliselect"
)

func RenderMenu() string {
	menu := gocliselect.NewMenu("Chose a option")
	menu.AddItem("Play", "play")
	menu.AddItem("Exit", "exit")
	return menu.Display()
	// width, height, _ := term.GetSize(0)
	// paddingX := (width - len("center")) / 2
	// paddingY := (height) / 2
	// fmt.Print(strings.Repeat("\n", paddingY))
	// fmt.Print(strings.Repeat(" ", paddingX))
	// fmt.Print("Center\n")
}
