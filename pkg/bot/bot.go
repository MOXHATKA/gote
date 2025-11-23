package bot

import (
	"context"
	"gote/pkg/api"
	"gote/pkg/types"
	"log"
)

type Config struct {
	Token          string
	Limit          int64
	Timeout        int64
	Offset         int64
	AllowedUpdates []string
}

type Bot struct {
	ctx          context.Context
	updateParams types.GetUpdates
	API          *api.API
	State        *StateStore
	Store        *Store
}

func NewBot(ctx context.Context, config Config) *Bot {
	return &Bot{
		ctx:   ctx,
		API:   api.NewAPI(config.Token),
		State: NewStateStore(),
		Store: NewStore(),
		updateParams: types.GetUpdates{
			Limit:   config.Limit,
			Timeout: config.Timeout,
			Offset:  config.Offset,
		},
	}
}

func (bot *Bot) Run() {
	for {
		select {
		case <-bot.ctx.Done():
			return
		default:
			response, err := bot.API.GetUpdates(bot.ctx, types.GetUpdates{
				Limit:   100,
				Timeout: 50,
				Offset:  bot.updateParams.Offset,
			})
			if err != nil {
				log.Println("Ошибка получения Update")
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

					state, ok := (*bot.State.States)[text]
					if ok {
						(*bot.State.UsersState)[id] = state
						state.Handle(ctx, &update, bot)
						continue
					}

					userState := (*bot.State.UsersState)[id]
					if userState != nil {
						userState.Handle(ctx, &update, bot)
					}
				}
			}(bot.ctx, response)

			lenUpdate := len(response)
			if lenUpdate > 0 {
				bot.updateParams.Offset = response[lenUpdate-1].UpdateId + 1
			}
		}
	}
}
