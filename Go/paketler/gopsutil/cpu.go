package main

import (
	"fmt"

	"github.com/raifpy/Go/errHandler"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	info, err := cpu.Info()
	errHandler.HandlerExit(err)
	//fmt.Println(info)
	Core := len(info)
	fmt.Println("Çekirdek sayısı = ", Core)

	persent, _ := cpu.Percent(5, true)            // Açıkçası anlayabilmiş değilim
	fmt.Println("Kullanım = %", persent[0], "\n") //
	ortalama(info)

	for core, cp := range info {
		// her çekirdek için yeni bir []
		fmt.Println("Core = ", core)
		fmt.Println("CPU     : ", cp.CPU)
		fmt.Println("VEN.ID  : ", cp.VendorID)
		fmt.Println("Family  : ", cp.Family)
		fmt.Println("Model   : ", cp.Model)
		fmt.Println("Steppng : ", cp.Stepping)
		fmt.Println("Phy.ID  : ", cp.PhysicalID)
		fmt.Println("CoreID  : ", cp.CoreID)
		fmt.Println("Cores   : ", cp.Cores)
		fmt.Println("ModlName: ", cp.ModelName)
		fmt.Println("Mhz     : ", cp.Mhz)
		fmt.Println()

		fmt.Println()
	}

}

func ortalama(info []cpu.InfoStat) {
	for core, cpu := range info {
		fmt.Println(core)
		fmt.Println(cpu)
	}

}
