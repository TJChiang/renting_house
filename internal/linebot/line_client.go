package linebot

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"os"
)

func Init() (*linebot.Client, error) {
	return linebot.New(
		os.Getenv("LINE_SECRET"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
}
