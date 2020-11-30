package main

import (
	"os"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/mem"

	"github.com/gizak/termui/v3/widgets"

	ui "github.com/gizak/termui/v3"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	//fmt.Println(getCPU())
	gui()

}

func gui() {
	ui.Init()
	defer func() {
		ui.Clear()
		ui.Close()
	}()
	en, boy := ui.TerminalDimensions()
	bar := widgets.NewBarChart()
	bar.SetRect(0, 0, en/2, boy)
	bar.Title = "Real-Time Cpu Mhz Bar"
	bar.BorderStyle.Fg = ui.ColorGreen
	bar.BarWidth = 10

	bar2 := widgets.NewGauge()
	bar2.Title = "Real-Time Mem Usage"
	bar2.BorderStyle.Fg = ui.ColorYellow
	bar2.SetRect(en/2, 0, en, boy/2)
	bar2.BarColor = ui.ColorGreen

	bar3 := widgets.NewGauge()
	bar3.Title = "Real-Time Swap Usage"
	bar3.BorderStyle.Fg = ui.ColorYellow
	bar3.SetRect(en/2, boy, en, boy/2)
	bar3.BarColor = ui.ColorGreen
	ui.Render(bar2, bar3)

	go func() {
		for {

			bar2.Percent = int(getMem().UsedPercent)
			if bar2.Percent > 50 && bar2.Percent < 70 {
				bar2.BarColor = ui.ColorYellow
			} else if bar2.Percent >= 70 {
				bar2.BarColor = ui.ColorRed
			} else {
				if bar2.BarColor != ui.ColorGreen {
					bar2.BarColor = ui.ColorGreen
				}
			}
			ui.Render(bar2)
			time.Sleep(time.Second)

		}
	}()

	go func() {
		for {
			bar3.Percent = int(getSwap().UsedPercent)
			if bar3.Percent > 50 && bar3.Percent < 70 && bar3.BarColor != ui.ColorYellow {
				bar3.BarColor = ui.ColorYellow
			} else if bar3.Percent >= 70 && bar3.BarColor != ui.ColorRed {
				bar3.BarColor = ui.ColorRed
			} else if bar3.Percent < 50 && bar3.BarColor != ui.ColorGreen {
				bar3.BarColor = ui.ColorGreen
			}
			ui.Render(bar3)
			time.Sleep(time.Second)
		}
	}()

	var sayi = 0

	fonk := func() {
		eleman := getCPU()
		if sayi == 0 {
			for i := 0; i < len(eleman); i++ {
				bar.Labels = append(bar.Labels, "Core "+strconv.Itoa(i))

			}
			bar.Data = make([]float64, len(eleman))

			//ui.Render(bar)

		}
		for i := 0; i < len(eleman); i++ {
			bar.Data[i] = eleman[i].Mhz
			ui.Render(bar)
		}
		sayi++

	}

	ui.Render(bar)

	go func() {
		for ch := range ui.PollEvents() {
			switch ch.ID {
			case "q", "<C-c>":
				ui.Clear()
				ui.Close()
				os.Exit(0)
			}
		}

	}()

	for {
		fonk()
		time.Sleep(time.Second)
		//fmt.Println("Sorun ne ulan")
	}

}

func getCPU() []cpu.InfoStat {
	e, err := cpu.Info()
	if err != nil {
		return []cpu.InfoStat{}
	}
	return e
}
func getMem() *mem.VirtualMemoryStat {
	e, err := mem.VirtualMemory()
	if err != nil {
		return nil
	}
	return e
}

func getSwap() *mem.SwapMemoryStat {
	e, _ := mem.SwapMemory()
	return e

}
