package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png" // Bu elemanları kullanmayacağız ama eklemek zorundayız :D
	"log"
	"net/http"

	"github.com/raifpy/Go/errHandler"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const url = "https://cdn4.telesco.pe/file/ucqUzuuo3oNTFQkBxRoANbvx3bgKi3yBq7pP57KoaSggaAOq1E7bppaNvBm86fQ9Nl4LHcVqoqzWx8Wg6CzPMnqhNkMhtnbz1h292cEHx6_9Uc9Dy4J0osRsn23ejmMvjl1ueAwXppdgWgnimXzkwHkywN-jdGbfJh1X0RDepgwW7kPxgyL178OMIo0QUk3OxqCJoOiwCaHwO5vRqQpkLCf3Cy4_aMIpaH0dmB4fauNZ-i5De5jUj-WRa16KzIcEUL9W3slaI70MZJM6IeiUQHxZQOD5Bdf9aRR7OI5TTdqj_W6QlxxnyRLrAv5w9kjOKxb82veBCmHsmzVQnM8ahg.jpg"

//const url = "https://media3.giphy.com/headers/colorful-courier/ammr2uN8JUE3.gif"

//const url = "https://teknobur.com/wp-content/uploads/2019/05/telegram.jpg"

func main() {
	hamImage, err := http.Get(url)
	if err != nil {
		log.Fatalln("image'e istek atarken hata ! : ", err.Error())
	}
	errHandler.HandlerExit(ui.Init())
	defer ui.Close()
	defer ui.Clear()

	resim, _, err := image.Decode(hamImage.Body)
	errHandler.HandlerExit(err)
	//os.Exit(1)

	resimUI := widgets.NewImage(resim)
	resimUI.Title = "NewImage(image)"
	//resm := resim.Bounds().Size()
	//en, boy := resm.Y, resm.X
	//resimUI.

	en, boy := ui.TerminalDimensions()
	resimUI.SetRect(0, 0, en, boy)
	//resimUI.SetRect(0, 0, 100, 40)

	text := widgets.NewParagraph()
	text.Title = "NewParagrapt"
	text.Text = "Merhaba KodWare !\nBu Program raifpy tarafından [Golang](fg:yellow) - Termiu kullanılarak yapılmıştır ."
	text.BorderStyle.Fg = ui.ColorYellow

	text.SetRect(130, 0, 20, 4) // Neden boyutlandırmama uymuyor anlayabilmiş değilim :)

	ui.Render(resimUI, text)

	fmt.Scanln()

}
