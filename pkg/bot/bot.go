package bot

import (
	"context"
	"gote/pkg/api"
	"gote/pkg/types"
	"log"
	"time"
)

type Config struct {
	Token           string
	Limit           int64
	Timeout         int64
	Offset          int64
	AllowedUpdates  []string
	UpdatesCapacity int64
}

type Bot struct {
	ctx             context.Context
	API             *api.API
	State           *StateStore
	Store           *Store
	Dependencies    *Dependencies
	updatesCapacity int64
	updateParams    types.GetUpdates
}

func NewBot(ctx context.Context, config Config) *Bot {
	if config.Limit <= 0 {
		config.Limit = 100
	}
	if config.Timeout <= 0 {
		config.Timeout = 30
	}
	if config.UpdatesCapacity <= 0 {
		config.UpdatesCapacity = 100
	}

	return &Bot{
		ctx:             ctx,
		API:             api.NewAPI(config.Token),
		State:           NewStateStore(),
		Store:           NewStore(),
		updatesCapacity: config.UpdatesCapacity,
		updateParams: types.GetUpdates{
			Limit:          config.Limit,
			Timeout:        config.Timeout,
			Offset:         config.Offset,
			AllowedUpdates: config.AllowedUpdates,
		},
	}
}

func (bot *Bot) AddDependencies(dd *Dependencies) {
	bot.Dependencies = dd
}

type UpdatesChannel <-chan types.Update

func (bot *Bot) GetUpdatesChannel() UpdatesChannel {
	ch := make(chan types.Update, bot.updatesCapacity)

	go func() {
		for {
			if bot.ctx.Err() != nil {
				close(ch)
				return
			}

			updates, err := bot.API.GetUpdates(bot.ctx, bot.updateParams)
			if err != nil {
				log.Println(err)
				log.Println("Ошибка получения обновлений, следующая попытка через 5 секунд...")
				time.Sleep(time.Second * 5)
				continue
			}

			for _, update := range updates {
				if update.UpdateId >= bot.updateParams.Offset {
					bot.updateParams.Offset = update.UpdateId + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func getChatID(u types.Update) (int64, bool) {
	var id int64
	founded := true

	switch {
	case u.Message != nil:
		id = u.Message.Chat.Id

	case u.EditedMessage != nil:
		id = u.EditedMessage.Chat.Id

	case u.ChannelPost != nil:
		id = u.ChannelPost.Chat.Id

	case u.EditedChannelPost != nil:
		id = u.EditedChannelPost.Chat.Id

	case u.BusinessMessage != nil:
		id = u.BusinessMessage.Chat.Id

	case u.EditedBusinessMessage != nil:
		id = u.EditedBusinessMessage.Chat.Id

	case u.DeletedBusinessMessages != nil:
		id = u.DeletedBusinessMessages.Chat.Id

	case u.CallbackQuery != nil && u.CallbackQuery.Message != nil:
		msg := (*u.CallbackQuery.Message)
		switch m := msg.(type) {
		case *types.Message:
			id = m.Chat.Id
		case *types.InaccessibleMessage:
			id = m.Chat.Id
		}

	case u.MyChatMember != nil:
		id = u.MyChatMember.Chat.Id

	case u.ChatMember != nil:
		id = u.ChatMember.Chat.Id

	case u.ChatJoinRequest != nil:
		id = u.ChatJoinRequest.Chat.Id

	case u.MessageReaction != nil:
		id = u.MessageReaction.Chat.Id

	case u.MessageReactionCount != nil:
		id = u.MessageReactionCount.Chat.Id

	case u.ChatBoost != nil:
		id = u.ChatBoost.Chat.Id

	case u.RemovedChatBoost != nil:
		id = u.RemovedChatBoost.Chat.Id

	default:
		founded = false
	}

	return id, founded
}
