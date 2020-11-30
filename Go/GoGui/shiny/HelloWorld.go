package main

import (
	"log"

	"golang.org/x/mobile/app"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	"golang.org/x/exp/shiny/widget/theme"
)

func oye() {
	driver.Main(func(s screen.Screen) {
		label := widget.NewLabel("Hello World!")
		/*button := newButton("Quit",
		func() {
			log.Println("To quit close this window")
		})*/

		w := widget.NewFlow(widget.AxisVertical, label)
		sheet := widget.NewSheet(widget.NewUniform(theme.Neutral, w))

		w.Measure(theme.Default, 0, 0)
		if err := widget.RunWindow(s, sheet, &widget.RunWindowOptions{
			NewWindowOptions: screen.NewWindowOptions{
				Title:  "Hello",
				Width:  w.MeasuredSize.X,
				Height: w.MeasuredSize.Y,
			},
		}); err != nil {
			log.Fatal(err)
		}
	})
}

func main() {
	app.Main(func(a app.App) {
		oye()
	})
}
