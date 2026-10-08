package ui

import "github.com/rivo/tview"

func (u *UI) buildRight() tview.Primitive {
	return newBox("Right (20 cols)")
}
