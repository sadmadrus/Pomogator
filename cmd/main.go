package main

import (
	"flag"
	"log"
	"pomogator/internal/bot"
)

func main() {
	token := mustToken
	b := bot.New(_, token)
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
