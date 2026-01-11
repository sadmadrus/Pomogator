package telegram

import "pomogator/internal/bot"

type Processor struct {
	tg     *bot.Client
	offset int
	//storage
}

func New(client bot.Client)
