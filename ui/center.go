package ui

import "github.com/rivo/tview"

func (u *UI) buildTop() tview.Primitive {
	return newBox("Top")
}

func (u *UI) buildMiddle() tview.Primitive {
	return newBox("Middle (3 x height of Top)")
}

func (u *UI) buildBottom() tview.Primitive {
	return newBox("Bottom (5 rows)")
}
