package main

import (
	"log"
	linebotClient "renting_house/internal/linebot"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	bot, err := linebotClient.Init()
	if err != nil {
		log.Fatalln("Initial linebot client error:", err)
	}

	bot.BroadcastMessage(linebot.NewTextMessage("Hi")).Do()
}
