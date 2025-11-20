package bot

import (
	"fmt"
	"gote/internal/commands"
	"gote/internal/handlers"
	"gote/internal/state"
	"gote/internal/utils/ctx"
	"gote/pkg/methods"
	"gote/pkg/types"
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
		ctx:      ctx,
		Commands: &commands.Commands{},
	}
	return bot
}

func (bot *Bot) OnCommand(command string, handler handlers.HandlerFunc) {
	(*bot.Commands)[command] = handler
}

// func (b *Bot) WithCommands(commands *commands.Commands) {
// 	b.Commands = commands
// }

// func (b *Bot) WithHandlers(handlers *handlers.Handlers) {
// 	b.Handlers = handlers
// }

func (b *Bot) WithState(stateMachine *state.StateMachine) {
	b.StateMachine = stateMachine
}

func (bot *Bot) RunUpdate() {
	for {
		select {
		case <-bot.ctx.GoContext.Done():
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

			go func(ctx ctx.CustomContext, updates []types.Update) {
				for _, update := range updates {
					msg := update.Message
					if msg == nil {
						continue
					}
					id := update.Message.Chat.Id

					text := msg.Text
					handlerFunc, ok := (*bot.Commands)[text]
					if ok {
						// bot.StateMachine.UsersState[id]
						// bot.StateMachine.UsersState[id].Action(ctx, &update)
						handlerFunc(ctx.GoContext, update)
						continue
					}

					if bot.StateMachine != nil {
						bot.StateMachine.NextState(ctx, &update)
					}
					fmt.Println(bot.StateMachine.GetState(id))
				}
			}(bot.ctx, response)

			lenUpdate := len(response)
			if lenUpdate > 0 {
				bot.offset = response[lenUpdate-1].UpdateId + 1
			}
		}
	}
}
