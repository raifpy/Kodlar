package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/raifpy/Go/errHandler"
)

func main() {
	errHandler.HandlerExit(ui.Init())
	defer ui.Close()

	eleman := widgets.NewParagraph()
	eleman.Title = "Merhaba Dünya"
	eleman.Text = `[Go](fg:green,bg:white) (diğer adıyla golang), Google'da 2007 yılından itibaren geliştirilmeye başlayan açık kaynak programlama dilidir.[4] Daha çok sistem programlama için tasarlanmış olup, derlenmiş ve statik tipli bir dildir. Kasım 2009'da çıkmıştır. Go derleyicisi "gc", açık kaynak yazılım olarak, Linux, OS X, Windows, bazı BSD ve Unix versiyonları, ve ayrıca 2015'ten itibaren akıllı telefonlar için geliştirilmiştir.`
	en, boy := ui.TerminalDimensions()
	eleman.SetRect(0, 0, en, boy)

	ui.Render(eleman)
	for abc := range ui.PollEvents() {
		/*if abc.Type == ui.KeyboardEvent {
			fmt.Println(abc)
		}*/
		fmt.Println(abc)
	}

	//time.Sleep(time.Second * 5)
	// [str](fg:color,bg:color) formatı ile renklendirebiliyoruz
}
