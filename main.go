package main

import (
	"log"

	"ASMG/ui"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	root := ui.BuildLayout()

	if err := app.SetRoot(root, true).SetFocus(root).Run(); err != nil {
		log.Fatal(err)
	}
}
