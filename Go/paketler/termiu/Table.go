//tablo1

package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	ui.Init()
	defer ui.Close()
	defer ui.Clear()

	en, boy := ui.TerminalDimensions()
	en, boy = en/2, boy/2

	tablo1 := widgets.NewTable()
	tablo1.Title = "tablo1"
	tablo1.Rows = [][]string{[]string{"Merhaba", "Dünya", "tablo1"}, []string{"tablo1", "Dünya", "Merhaba"}}
	tablo1.TextAlignment = ui.AlignCenter

	tablo1.SetRect(en/2, 0, en, boy)

	tablo2 := widgets.NewTable()
	tablo2.Title = "tablo2"
	//tablo2.BorderBottom = true // Birşey değişmedi ama mutlaka işe yarıyordur
	tablo2.Rows = [][]string{[]string{"Merhaba", "Dünya", "tablo2"}, []string{"tablo2", "Dünya", "Merhaba"}}
	tablo2.TextAlignment = ui.AlignLeft
	tablo2.TextStyle.Fg = ui.ColorYellow
	tablo2.RowSeparator = false // satur ayırıcısını kapat
	tablo2.SetRect(0, 0, en/2, boy)

	tablo3 := widgets.NewTable()
	tablo3.Title = "tablo3"
	tablo3.Rows = [][]string{[]string{"Merhaba", "Dünya", "tablo3"}, []string{"tablo3", "Dünya", "Merhaba"}}
	//tablo3.BorderRight = true
	tablo3.SetRect(120, 0, en, boy) // Bunu ayarlayamadım :D
	tablo3.TextAlignment = ui.AlignRight
	tablo3.TextStyle.Fg = ui.ColorWhite

	tablo3.FillRow = true                            // Bu eleman ile satırın tamamını renklendirebiliyoruz ama olmadı :D
	tablo3.RowStyles[0] = ui.NewStyle(ui.ColorGreen) // Bu şekilde sadece bir satırı renklendirebiliriz

	ui.Render(tablo1, tablo2, tablo3)

	fmt.Scanln()

}
