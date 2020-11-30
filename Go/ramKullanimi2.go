package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func gc() {
	runtime.GC()
}

func getUsage() string {
	var abc runtime.MemStats
	runtime.ReadMemStats(&abc)
	return time.Now().String()[11:22] + "   " + strconv.Itoa(int(abc.Alloc/1024/1024)) + " mb"
}

func main() {
	fmt.Println(getUsage())
	time.Sleep(time.Second * 3)
	for {
		go func() {
			//fmt.Sprintln("Merhaba Dünya")
			fmt.Println(getUsage())
			gc()
			for {
				v := make([]string, 500)
				fmt.Sprintln(v)
			}
		}()
		time.Sleep(time.Second * 3)
	}
}
