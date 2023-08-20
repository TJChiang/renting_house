package main

import (
	"fmt"
	"log"
	"renting_house/services/house591"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
		return
	}

	// 爬取資料
	crawler, err := house591.Default()
	if err != nil {
		log.Fatalln("Initial 591 crawler failed:", err)
		return
	}
	data, err := crawler.FetchHouses(house591.DefaultOptions())
	if err != nil {
		log.Fatalln("Fetch data failed:", err)
		return
	}
	fmt.Println("Status:", data.Status)

	// Line Message API
	// bot, err := line_message.Init()
	// if err != nil {
	// 	log.Fatalln("Initial linebot client error:", err)
	// 	return
	// }

	// contents, err := linebot.UnmarshalFlexMessageJSON()
	// if err != nil {
	// 	log.Fatalln("Parse json flex message error: ", err)
	// 	return
	// }

	// bot.BroadcastMessage(
	// 	linebot.NewFlexMessage("carousel", &linebot.CarouselContainer{}),
	// ).Do()
}
