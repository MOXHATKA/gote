package main

import (
	"context"
	"fmt"
	"gote/internal/env"
	gb "gote/pkg/bot"
	"os"
)

type MyService struct {
	Name string
}

func (ms MyService) SayHello() {
	fmt.Println("Hello,", ms.Name)
}

func main() {
	// получение ключа бота из файла .env
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	// создание контекста
	ctx, closeFunc := context.WithCancel(context.Background())
	defer closeFunc()

	// создание бота
	bot := gb.NewBot(ctx, gb.Config{
		Token:   token,
		Limit:   100,
		Timeout: 30,
	})

	// добавление зависимостей (DI)
	deps := gb.NewDependencies()

	myService := MyService{"My Services"}
	deps.Provide(myService)

	bot.AddDependencies(deps)

	// создание состояний
	startState := bot.State.NewState("start", RequestName)
	_ = bot.State.NewState("writeNameState", WriteName)
	_ = bot.State.NewState("requestMailState", RequestMail)
	_ = bot.State.NewState("writeMailState", WriteMail)

	// рабора с командами
	bot.State.OnCommand("/start", startState)

	updates := bot.GetUpdatesChannel()
	for u := range updates {
		fmt.Println(u.Message.Text)
	}

}
