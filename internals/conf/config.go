package conf

import (
	"os"
	"strconv"

	"github.com/MukundSinghRajput/InlinePinterestBot/internals/logger"
	_ "github.com/joho/godotenv/autoload"
)

type configs struct {
	Token    string
	DB_URI   string
	OWNER_ID int64
}

var Config configs

func init() {
	log := logger.NewLogger("CONFIG")

	Owner, err := strconv.Atoi(os.Getenv("OWNER_ID"))
	if err != nil {
		log.Error("OWNER_ID should be a number")
		os.Exit(1)
	}

	Config = configs{
		Token:    os.Getenv("BOT_TOKEN"),
		DB_URI:   os.Getenv("DB_URI"),
		OWNER_ID: int64(Owner),
	}

	log.Info("Loaded Configs")
}
