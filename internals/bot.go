package internals

import (
	"os"
	"time"

	"github.com/MukundSinghRajput/InlinePinterestBot/internals/conf"
	"github.com/MukundSinghRajput/InlinePinterestBot/internals/loader"
	"github.com/MukundSinghRajput/InlinePinterestBot/internals/logger"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Bot struct {
	Token string
}

func NewBot() *Bot {
	bot := Bot{
		Token: conf.Config.Token,
	}
	return &bot
}

func (b *Bot) Start() {
	log := logger.NewLogger("BOT")
	if b.Token == "" || len(b.Token) != 46 {
		log.Error("Please entere correct token")
		os.Exit(1)
	}
	bot, err := gotgbot.NewBot(b.Token, &gotgbot.BotOpts{
		// DisableTokenCheck: true,
	})

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			b.SendMessage(conf.Config.OWNER_ID, err.Error(), &gotgbot.SendMessageOpts{})
			return ext.DispatcherActionNoop
		},
		MaxRoutines: -1,
	})

	updater := ext.NewUpdater(dispatcher, nil)
	loader.LoadMethods(dispatcher)

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates:    true,
		EnableWebhookDeletion: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})

	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Infof("%s has started....", bot.Username)
	updater.Idle()
}
