package methods

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	message := "Hello *%s* there i am *Pinterest* bot to send images from pinterest to your friends instantly in the chat"
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf(message, ctx.EffectiveUser.FirstName), &gotgbot.SendMessageOpts{
		ParseMode: "MARKDOWN",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{
					{Text: "Channel", Url: "https://t.me/MukundX"},
				},
			},
		},
	})
	return err
}
