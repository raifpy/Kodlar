package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {

	window := app.NewWindow()
	for e := range window.Events() {
		gtx := layout.NewContext(&op.Ops{}, e)
		l := material.H1(window, "hey")
		l.Layout(gtx)
		e.Frame(gtx.Ops)
	}

}
