package main

import (
	"context"
	"fmt"
	"gote/internal/bot"
	"gote/pkg/methods"
	"gote/pkg/types"
	c "gote/internal/commands"
	h "gote/internal/handlers"
	"gote/internal/utils/ctx"
	"gote/internal/utils/env"
	"os"
)

func main() {
	_ = env.Load(".env")
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Токен отсутствует")
	}

	ctx := ctx.CustomContext{
		Token: token,
		GoContext: context.Background(),
	}

	b := bot.NewBot(ctx)

	me, _ := methods.GetMe(ctx, types.GetMe{})
	fmt.Println(me.FirstName, me.Id, me.Username)
	
	commands := c.NewCommands()
	commands.Add("/start", StartHandler)
	b.WithCommands(&commands)

	handlers := h.NewHandlers()
	// handlers.Add(h.Message, MessageHandler)
	b.WithHandlers(&handlers)

	stateMachine := createStateMachine()

	b.WithState(stateMachine)

	b.RunUpdate()
}
