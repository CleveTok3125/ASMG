package ui

import "github.com/rivo/tview"

func newBox(title string) *tview.Box {
	return tview.NewBox().
		SetBorder(true).
		SetTitle(title)
}
