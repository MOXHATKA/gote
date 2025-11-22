package main

import (
	"context"
	gotebot "gote/internal/bot"
	"gote/internal/utils/env"
	"os"
)

const (
	startState       = "start"
	writeNameState   = "write_name"
	requestMailState = "request_mail"
	writeMailState   = "write_mail"
)

func main() {
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	ctxClose, closeFunc := context.WithCancel(context.Background())
	ctx := context.WithValue(ctxClose, "token", token)

	bot := gotebot.NewBot(ctx)

	bot.OnState(startState, RequestName)
	bot.OnState(writeNameState, WriteName)
	bot.OnState(requestMailState, RequestMail)
	bot.OnState(writeMailState, WriteMail)

	bot.OnCommand("/start", startState)

	bot.RunUpdate()
	closeFunc()
}
