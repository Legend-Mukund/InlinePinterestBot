package loader

import (
	"github.com/MukundSinghRajput/InlinePinterestBot/internals/methods"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/inlinequery"
)

func LoadMethods(dis *ext.Dispatcher) {
	dis.AddHandler(handlers.NewCommand("start", methods.Start))
	dis.AddHandler(handlers.NewInlineQuery(inlinequery.All, methods.Pin))
	// dis.AddHandler(handlers.NewCommand("source", modules.Source))
	// dis.AddHandler(handlers.NewCommand("repo", modules.Source))
}
