package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	adaptor := firmata.NewAdaptor("/dev/ttyUSB0") // ttyUSB0 kısmı bana özel.

	gomuluLed := gpio.NewLedDriver(adaptor, "13")
	pin8 := gpio.NewLedDriver(adaptor, "8")

	work := func() {
		for {
			time.Sleep(time.Second * 3)
			pin8.On()
			gomuluLed.Off()

			time.Sleep(time.Second / 10)
			pin8.Off()
			gomuluLed.On()
		}
	}

	bot := gobot.NewRobot("bot",
		[]gobot.Device{gomuluLed},
		[]gobot.Connection{adaptor},
		work)

	bot.Start()

}
