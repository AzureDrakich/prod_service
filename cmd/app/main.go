package main

import (
	"app/internal/app"
	"app/internal/config"
	"log"
)

func main() {
	log.Print("config init")
	cfg := config.GetConfig()
	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("application started")
	a.Run()
}
