package main

import (
	"log"

	"github.com/brunobriante/rodanodeck_bot/internal/bot"
)

func main() {
	b, err := bot.New()
	if err != nil {
		log.Fatal(err)
	}

	bot.Start(b)
}
