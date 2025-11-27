package main

import (
	"context"
	"gote/internal/env"
	"gote/pkg/core"
	"gote/pkg/types"
	"gote/pkg/updater"
	"os"
)

func main() {
	// получение ключа бота из файла .env
	// BOT_TOKEN=токен_из_BotFather
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	// создание контекста
	ctx, closeFunc := context.WithCancel(context.Background())
	defer closeFunc()

	// создание бота
	b := core.NewBot(ctx, token,
		core.WithHTTPClient(nil),
		core.WithLogger(nil),
	)

	// полдучение обновлений от Telegram
	poller := updater.NewPoller(b)
	updates := poller.Start()
	for u := range updates {
		if cb := u.CallbackQuery; cb != nil {
			message, err := types.CastTo[types.Message](cb.Message)
			if err != nil {
				b.Logger().Error(err.Error())
			}

			b.Logger().Info(message.Text)
			// b.SendMessage(ctx, types.SendMessage{
			// 	ChatId: u.CallbackQuery.Message),
			// 	Text:   "text " + u.CallbackQuery.Data,
			// },

		}

		if msg := u.Message; msg != nil {
			buttons := []types.InlineKeyboardButton{
				{
					Text:         "1",
					CallbackData: "1",
				},
				{
					Text:         "2",
					CallbackData: "2",
				},
			}

			b.SendMessage(ctx, types.SendMessage{
				ChatId: u.Message.Chat.Id,
				Text:   msg.Text,
				ReplyMarkup: types.InlineKeyboardMarkup{
					InlineKeyboard: [][]types.InlineKeyboardButton{
						buttons,
					},
				},
			})
		}
	}
}
