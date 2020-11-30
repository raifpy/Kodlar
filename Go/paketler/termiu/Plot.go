// şablon biçiminde birşeyler yapacak :)

package main

import (
	"math/rand"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/raifpy/Go/errHandler"
)

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(max-min+1) + min)
}

func main() {
	errHandler.HandlerExit(ui.Init())
	defer func() {
		ui.Clear()
		ui.Close()
	}()

	defer func() {
		abc := recover()
		if abc != nil {
			main()
		}
	}()
	plot := widgets.NewPlot()
	plot.Title = "NewPlot"
	plot.BorderStyle.Fg = ui.ColorYellow
	plot.TitleStyle.Fg = ui.ColorGreen
	plot.AxesColor = ui.ColorRed      // eksen rengi imiş ama farkı göremedim
	plot.LineColors[0] = ui.ColorBlue // çizgi elemanın rengi
	//plot.Marker = widgets.MarkerDot
	plot.Marker = widgets.MarkerBraille // default bu , dot'dan çok daha iyi

	floatslice := make([][]float64, 0)

	for i := 0; i < randomInt(0, 5); i++ {
		listeInt := randomInt(1, 10)
		//log.Panicln(listeInt)
		liste := make([]float64, listeInt)

		for ii := 0; ii < listeInt; ii++ {

			liste[ii] = float64(randomInt(0, 3))

		}

		floatslice = append(floatslice, liste)

	}

	/*plot.Data = [][]float64{
		[]float64{1, 2, 3, 5, 6, 7, 5, 10, 3},
		[]float64{2, 2, 4, 8, 7, 8, 7, 8, 6}, // istediğin gibi ekleyebilirisin . LineColors[index]  ile de reng belirle ..
	}*/
	plot.Data = floatslice

	en, boy := ui.TerminalDimensions()
	plot.SetRect(0, 0, en, boy)
	ui.Render(plot)
	//log.Println(plot.Data)

	for ch := range ui.PollEvents() {
		switch ch.ID {
		case "q", "<C-c>":
			return
		case "u", "<Up>":
			func() {
				for indexBir, floatList := range plot.Data {
					for indexIki, list := range floatList {
						plot.Data[indexBir][indexIki] = list + 1
					}
				}
				ui.Render(plot)

			}()

		case "d", "<Down>":
			func() {
				for indexBir, floatList := range plot.Data {
					for indexIki, list := range floatList {
						plot.Data[indexBir][indexIki] = list - 1
					}
				}

				ui.Render(plot)

			}()

		case "<Right>":
			func() {
				for indexBir, floatList := range plot.Data {
					for indexIki, list := range floatList {
						plot.Data[indexBir][indexIki] = list * 10
					}
				}

				ui.Render(plot)

			}()
		case "<Left>":
			func() {
				for indexBir, floatList := range plot.Data {
					for indexIki, list := range floatList {
						plot.Data[indexBir][indexIki] = list / 10
					}
				}

				ui.Render(plot)

			}()
		}
	}

}
