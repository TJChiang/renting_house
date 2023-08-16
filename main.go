package main

import (
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	lineclient "renting_house/internal/linebot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	bot, err := lineclient.Init()

	if err != nil {
		log.Fatalln(err)
	}

	bot.PushMessage("", linebot.NewTextMessage("Hi")).Do()
}
