package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
	"time"
)

func main() {
	getCpu()
}

func getCpu() {
	for range time.Tick(time.Second * 1) {
		percent, err := cpu.Percent(1*time.Second, false)
		if err != nil {
			log.Fatal(err)
		}

		cpuString := fmt.Sprintf("%.2f", percent[0])
		fmt.Println(cpuString, "%")
	}
}
