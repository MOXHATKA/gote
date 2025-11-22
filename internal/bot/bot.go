package bot

import (
	"context"
	"fmt"
	"gote/pkg/methods"
	"gote/pkg/types"
	"log"
)

type Bot struct {
	ctx        context.Context
	offset     int64
	States     *States
	UsersState *UsersState
}

func NewBot(ctx context.Context) *Bot {
	bot := &Bot{
		ctx:        ctx,
		States:     &States{},
		UsersState: &UsersState{},
	}
	return bot
}

func (bot *Bot) OnCommand(command string, stateName string) {
	(*bot.States)[command] = (*bot.States)[stateName]
}

func (bot *Bot) RunUpdate() {
	for {
		select {
		case <-bot.ctx.Done():
			return
		default:
			response, err := methods.GetUpdates(bot.ctx, types.GetUpdates{
				Limit:   100,
				Timeout: 50,
				Offset:  bot.offset,
			})
			if err != nil {
				log.Println("Не получилось получить Update")
				return
			}

			go func(ctx context.Context, updates []types.Update) {
				for _, update := range updates {
					msg := update.Message
					if msg == nil {
						continue
					}

					id := msg.Chat.Id
					text := msg.Text

					state, ok := (*bot.States)[text]
					if ok {
						(*bot.UsersState)[id] = state
						state.Handle(ctx, &update, bot)
						continue
					}

					userState := (*bot.UsersState)[id]
					if userState != nil {
						fmt.Println(userState.Name)
						userState.Handle(ctx, &update, bot)
					}
				}
			}(bot.ctx, response)

			lenUpdate := len(response)
			if lenUpdate > 0 {
				bot.offset = response[lenUpdate-1].UpdateId + 1
			}
		}
	}
}
