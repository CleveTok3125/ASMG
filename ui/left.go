package ui

import "github.com/rivo/tview"

func (u *UI) buildLeft() tview.Primitive {
	return newBox("Left (1/2 x width of Top)")
}
