package main

import (
	"Ejemplo1/Controller"
	"Ejemplo1/Database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"log"
	"time"
)

func main() {
	app := fiber.New()
	if err := Database.Connect(); err != nil {
		log.Fatalln(err)
	}

	// Borrado de datos en Mongo

	go getCpu()
	go getStorage()

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
		Controller.InsertData("cpu", cpuString)
	}
}

func getStorage() {
	for range time.Tick(time.Second * 1) {
		percent, err := disk.Usage("/home")
		if err != nil {
			log.Fatalln(err)
		}

		storage := percent.UsedPercent
		storageString := fmt.Sprintf("%.2f", storage)
		fmt.Println(storageString, "%")
		Controller.InsertData("storage", storageString)
	}
}
