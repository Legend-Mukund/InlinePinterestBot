package methods

import (
	"fmt"
	"math/rand"
	"slices"

	"github.com/MukundSinghRajput/InlinePinterestBot/internals/api"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Pin(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.InlineQuery.Query
	emptyText := " "

	cm, err := b.GetChatMember(-1002146498946, ctx.InlineQuery.From.Id, &gotgbot.GetChatMemberOpts{})

	if err != nil {
		_, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
			gotgbot.InlineQueryResultArticle{
				Id:    fmt.Sprintf("%d", rand.Int()),
				Title: "Errorr !!!",
				InputMessageContent: gotgbot.InputTextMessageContent{
					MessageText: fmt.Sprintf("An error occured: %v", err),
				},
				Description: "An error occured",
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "Search", SwitchInlineQuery: &emptyText},
					}},
				},
			},
		}, &gotgbot.AnswerInlineQueryOpts{
			CacheTime:  0,
			IsPersonal: true,
		})
		return err
	}

	if cm.GetStatus() == "kicked" {
		_, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
			gotgbot.InlineQueryResultArticle{
				Id:    fmt.Sprintf("%d", rand.Int()),
				Title: "Banned !!!",
				InputMessageContent: gotgbot.InputTextMessageContent{
					MessageText: "Sadly you have been banned from @MukundX therefore you can't use me.",
				},
				Description: "You have been banned from @MukundX",
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "Report", Url: "https://t.me/MukundSRajput?text=Hey+I+have+been+banned+from+your+channel+@MukundX+look+into+this+maybe+it's+just+a+mistake"},
					}},
				},
			},
		}, &gotgbot.AnswerInlineQueryOpts{
			CacheTime:  0,
			IsPersonal: true,
		})
		return err
	}

	x := []string{"creator", "administrator", "member"}

	if !slices.Contains(x, cm.GetStatus()) {
		_, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
			gotgbot.InlineQueryResultArticle{
				Id:    fmt.Sprintf("%d", rand.Int()),
				Title: "Ehhhh !!!",
				InputMessageContent: gotgbot.InputTextMessageContent{
					MessageText: "Join @MukundX to use me",
				},
				Description: "Join @MukundX first",
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "Join", Url: "https://t.me/MukundX"},
					}},
				},
			},
		}, &gotgbot.AnswerInlineQueryOpts{
			CacheTime:  0,
			IsPersonal: true,
		})
		return err
	}

	if query == "" {
		_, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
			gotgbot.InlineQueryResultArticle{
				Id:    fmt.Sprintf("%d", rand.Int()),
				Title: "No query",
				InputMessageContent: gotgbot.InputTextMessageContent{
					MessageText: "Please enter a query to search for",
				},
				Description: "Please enter a query to search for",
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "Search", SwitchInlineQuery: &emptyText},
					}},
				},
			},
		}, &gotgbot.AnswerInlineQueryOpts{
			CacheTime:  0,
			IsPersonal: true,
		})
		return err
	}

	p := api.NewPinterest()
	images, err := p.Scrap(query, 40)
	if err != nil {
		return err
	}

	if len(images) == 0 {
		_, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
			gotgbot.InlineQueryResultArticle{
				Id:    fmt.Sprintf("%d", rand.Int()),
				Title: "Not Found",
				InputMessageContent: gotgbot.InputTextMessageContent{
					ParseMode:   "MARKDOWN",
					MessageText: fmt.Sprintf("The query *%s* wasn't found", query),
					LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
						IsDisabled: true,
					},
				},
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "Search", SwitchInlineQuery: &emptyText},
					}},
				},
				Description: fmt.Sprintf("The query %s wasn't found", query),
			},
		}, &gotgbot.AnswerInlineQueryOpts{
			IsPersonal: true,
			CacheTime:  0,
		})
		if err != nil {
			return err
		}
		return nil
	}

	var articles []gotgbot.InlineQueryResult
	for i, imageURL := range images {
		articles = append(articles, gotgbot.InlineQueryResultPhoto{
			Id:           fmt.Sprintf("%d", i),
			PhotoUrl:     imageURL,
			ThumbnailUrl: imageURL,
			Title:        fmt.Sprintf("Image %d", i+1),
			Description:  fmt.Sprintf("Pinterest Image %d", i+1),
			ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Search Again", SwitchInlineQuery: &emptyText},
				}},
			},
		})
	}

	_, err = ctx.InlineQuery.Answer(b, articles, &gotgbot.AnswerInlineQueryOpts{})

	if err != nil {
		_, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
			gotgbot.InlineQueryResultArticle{
				Id:    fmt.Sprintf("%d", rand.Int()),
				Title: "Errorr !!!",
				InputMessageContent: gotgbot.InputTextMessageContent{
					MessageText: fmt.Sprintf("Some error occured while fetching the images while searching for %v", query),
				},
				Description: "Some error occured while fetching the images",
				ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "Search", SwitchInlineQuery: &emptyText},
					}},
				},
			},
		}, &gotgbot.AnswerInlineQueryOpts{
			CacheTime: 0,
		})
		return err
	}
	return nil
}
