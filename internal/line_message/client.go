package line_message

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineMessage struct {
	Client *linebot.Client
}

func Init() (*LineMessage, error) {
	client, err := linebot.New(
		os.Getenv("LINE_SECRET"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
	if err != nil {
		return nil, err
	}

	return &LineMessage{
		Client: client,
	}, nil
}
