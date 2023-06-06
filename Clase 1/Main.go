package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"log"
	"time"
)

func main() {
	app := fiber.New()

	go getCpu()

	err := app.Listen(":8000")
	if err != nil {
		log.Fatalln("Error ", err)
	}

	time.Sleep(time.Second * 500)
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
