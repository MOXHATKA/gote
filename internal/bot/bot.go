package bot

import (
	"context"
	"fmt"
	"gote/pkg/types"
	"log"
)

type Bot struct {
	Token        string
	ctx          context.Context
	offset       int64
	Commands     map[string]Handler
	Handlers     map[string]Handler
	StateMachine StateMachine
}

func (b *Bot) AddCommand(text string, handler Handler) {
	b.Commands[text] = handler
}

type Handler func(context.Context, types.Update)

func NewBot(token string) *Bot {
	startState := State{
		Name:      "Старт",
		Condition: "/start",
	}

	endState := State{
		Name:      "Конец",
		Condition: "/end",
	}

	menuState := State{
		Name:      "Меню",
		Condition: "/menu",
	}

	startState.Children = append(startState.Children, &endState)
	endState.Children = append(endState.Children, &menuState)
	menuState.Children = append(menuState.Children, []*State{
		&startState,
		&endState,
	}...)

	bot := &Bot{
		Token: token,
		ctx:   context.Background(),
	}
	bot.Commands = map[string]Handler{}
	bot.Handlers = map[string]Handler{}
	bot.StateMachine = StateMachine{
		States:     map[int64]*State{},
		StartState: &startState,
		ResetState: &menuState,
	}
	return bot
}

func (bot *Bot) RunUpdate() {
	for {
		select {
		case <-bot.ctx.Done():
			return
		default:
			response, err := bot.GetUpdates(bot.ctx, types.GetUpdates{
				Limit:   100,
				Timeout: 50,
				Offset:  bot.offset,
			})
			if err != nil {
				log.Println("Чота пошло не так")
				return
			}

			go func(ctx context.Context, updates []types.Update) {
				for _, update := range updates {
					id := update.Message.Chat.Id
					bot.StateMachine.SetState(&update)
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
