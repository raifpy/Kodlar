package main

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	memm, _ := mem.VirtualMemory()
	total := memm.Total
	used := memm.Used
	oran := memm.UsedPercent
	fmt.Println(used, " / ", total, " :  %", int(oran))
	fmt.Println(sekilVer(float64(used)), "/", sekilVer(float64(total)), " :  %", int(oran))
}

func sekilVer(nesne float64) string {
	for _, birim := range []string{"", "K", "M", "G", "T", "P"} {
		if nesne < 1024 {
			if nesneStr := strings.Split(fmt.Sprint(nesne), "."); len(nesneStr) == 2 {
				ek := nesneStr[1]
				if len(ek) > 2 {
					return nesneStr[0] + "." + ek[:2] + birim
				} else {
					return nesneStr[0] + "." + ek + birim
				}
			}
			//return fmt.Sprintf("%v %s", nesne, birim)
		}
		nesne /= 1024
	}
	return ""
}
