package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"os"
	"renting_house/internal/line_message"
	"renting_house/services"
	"renting_house/services/house591"
	"strconv"
	"time"
)

func main() {
	// 載入 env
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
	bot, err := line_message.Init()
	if err != nil {
		log.Fatalln("Initial linebot client error:", err)
		return
	}

	carouselLimit, err := strconv.Atoi(os.Getenv("LINE_CAROUSEL_LIMIT"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 整理資料
	bubbleList := make([]*linebot.BubbleContainer, 0, carouselLimit)
	timezone, _ := time.LoadLocation("Asia/Taipei")
	datetime := time.Now().In(timezone).Format("2006-02-01 15:04:05")
	bot.Client.BroadcastMessage(linebot.NewTextMessage("[" + datetime + "] New house items.")).Do()
	for _, v := range data.Data.Data {
		log.Println("post id: ", v.PostId)
		bubbleList = append(bubbleList, services.GetBubbleContainer(v))
		// bubble container 滿了之後，放到 carousel container ，並清除
		if len(bubbleList) >= carouselLimit {
			// 發布租屋資料
			res, err := broadcastMessage(bot.Client, bubbleList)
			if err != nil {
				log.Fatalln("Response: ", res)
				log.Fatalln("Error: ", err)
				return
			}
			bubbleList = bubbleList[:0]
			time.Sleep(500 * time.Millisecond)
		}
	}

	// 把 list 剩下的資料發布出去
	if len(bubbleList) != 0 {
		res, err := broadcastMessage(bot.Client, bubbleList)
		if err != nil {
			log.Fatalln("Response: ", res)
			log.Fatalln("Error: ", err)
			return
		}
		bubbleList = bubbleList[:0]
	}
}

func broadcastMessage(bot *linebot.Client, bubbleList []*linebot.BubbleContainer) (*linebot.BasicResponse, error) {
	return bot.BroadcastMessage(linebot.NewFlexMessage("carousel", &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: bubbleList,
	})).Do()
}
