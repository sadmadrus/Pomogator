package main

import (
	"flag"
	"log"
	"pomogator/internal/bot"
)

const (
	tg_host = "api.telegram.org"
)

func main() {
	b := bot.New(tg_host, mustToken())
}

func mustToken() string {
	tok := flag.String(
		"bot_token",
		"",
		"Telegram bot token")

	flag.Parse()

	if *tok == "" {
		log.Fatal("token is not specified")
	}

	return *tok
}
