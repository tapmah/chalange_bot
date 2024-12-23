package main

import (
	"log"

	"github.com/tapmah/chalange_bot/internal/bot"
	"github.com/tapmah/chalange_bot/internal/config"
	"github.com/tapmah/chalange_bot/internal/db"
)

func init() {
	config.LoadConfig()
	db.InitDB()
}

func main() {
	cfg := config.GetConfig()

	tgBot, err := bot.NewTelegramBot(cfg.Token)
	if err != nil {
		log.Fatalf("could not create bot: %v", err)
	}

	tgBot.Start()
}
