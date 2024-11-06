package methods

import (
	"fmt"

	"github.com/MukundSinghRajput/InlinePinterestBot/internals/api"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Pin(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.InlineQuery.Query
	if query == "" {
		return nil
	}

	p := api.NewPinterest()
	images, err := p.Scrap(query, 40)
	if err != nil {
		return err
	}

	var articles []gotgbot.InlineQueryResult
	for i, imageURL := range images {
		articles = append(articles, gotgbot.InlineQueryResultPhoto{
			Id:           fmt.Sprintf("%d", i),
			PhotoUrl:     imageURL,
			ThumbnailUrl: imageURL,
			Title:        fmt.Sprintf("Image %d", i+1),
			Description:  fmt.Sprintf("Pinterest Image %d", i+1),
		})
	}

	_, err = ctx.InlineQuery.Answer(b, articles, &gotgbot.AnswerInlineQueryOpts{
		IsPersonal: true,
	})
	if err != nil {
		return fmt.Errorf("failed to send inline query results: %w", err)
	}
	return nil
}
