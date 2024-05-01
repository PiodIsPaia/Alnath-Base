package main

import (
	"log"

	"github.com/piodispaia/alnath/internal/bot"
	"github.com/piodispaia/alnath/internal/settings"
)

func main() {
	config, err := settings.GetConfig()
	if err != nil {
		log.Fatal("Error getting configuration:", err)
	}

	bot.Start(config.Core.Token)
}
