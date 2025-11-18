package main

import (
	"context"
	"fmt"
	"gote/internal/bot"
	"gote/internal/utils/env"
	"gote/pkg/types"
	"os"
)

func StartHandler(ctx context.Context, update types.Update) {
	fmt.Println("Я сказала стартуем!")
}

func main() {
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	bot := bot.NewBot(token)
	bot.AddCommand("/start", StartHandler)
	ctx := context.Background()
	user, err := bot.GetMe(ctx, types.GetMe{})
	if err != nil {
		fmt.Println("Ошибка GetMe:", err)
	}

	fmt.Printf("Бот: %s (@%s), ID: %d\n", user.FirstName, user.Username, user.Id)

	bot.RunUpdate()
}
