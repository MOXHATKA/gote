package main

import (
	"context"
	"gote/internal/env"
	gotebot "gote/pkg/bot"
	"log"
	"os"
)

func main() {
	// получение ключа бота из файла .env
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	// создание контекста
	ctx, close := context.WithCancel(context.Background())
	defer close()

	// создание бота
	bot := gotebot.NewBot(ctx, gotebot.Config{
		Token:   token,
		Limit:   100,
		Timeout: 50,
	})

	// создание состояний
	startState := bot.State.NewState("start", RequestName)
	_ = bot.State.NewState("writeNameState", WriteName)
	_ = bot.State.NewState("requestMailState", RequestMail)
	_ = bot.State.NewState("writeMailState", WriteMail)

	// рабора с командами
	bot.State.OnCommand("/start", startState)

	// запуск цикла обновлений
	err := bot.Run()
	if err != nil {
		log.Println(err)
	}
}
