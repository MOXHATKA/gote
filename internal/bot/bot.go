package bot

import (
	"context"
	"fmt"
	"gote/internal/commands"
	"gote/internal/handlers"
	"gote/internal/state"
	"gote/pkg/methods"
	"gote/pkg/types"
	"gote/internal/utils/ctx"
	"log"
)

type Bot struct {
	ctx          ctx.CustomContext
	offset       int64
	Commands     *commands.Commands
	Handlers     *handlers.Handlers
	StateMachine *state.StateMachine
	// enabledModules 	 []string
}

func NewBot(ctx ctx.CustomContext) *Bot {
	bot := &Bot{
		ctx: ctx,
	}
	return bot
}

func (b *Bot) WithCommands(commands *commands.Commands) {
	b.Commands = commands
	// b.enabledModules = append(b.enabledModules, "Commands")
}

func (b *Bot) WithHandlers(handlers *handlers.Handlers) {
	b.Handlers = handlers
	// b.enabledModules = append(b.enabledModules, "Handlers")
}

func (b *Bot) WithState(stateMachine *state.StateMachine) {
	b.StateMachine = stateMachine
	// b.enabledModules = append(b.enabledModules, "StateMachine")
}

func (bot *Bot) RunUpdate() {
	for {
		select {
		case <- bot.ctx.GoContext.Done():
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
					id := update.Message.Chat.Id

					text := msg.Text
					if string(text[0]) == "/" {
						handlerFunc, ok := (*bot.Commands)[text]
						if ok {
							handlerFunc(ctx, update)
							continue
						}
					}
					if bot.StateMachine != nil {
						bot.StateMachine.SetState(&update)
					}
					fmt.Println(bot.StateMachine.GetState(id))
				}
			}(bot.ctx.GoContext, response)

			lenUpdate := len(response)
			if lenUpdate > 0 {
				bot.offset = response[lenUpdate-1].UpdateId + 1
			}
		}
	}
}
