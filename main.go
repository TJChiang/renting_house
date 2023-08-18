package main

import (
	"fmt"
	"log"
	linebotClient "renting_house/internal/linebot"
	"renting_house/services/house591"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	// s := string(`{
	// 	"status": 1,
	// 	"data": {
	// 		"topData": [
	// 			{
	// 				"title": "\ud83c\udf1f\u7368\u6d17\u66ec\ud83c\udf1f\u7cfb\u7d71\u6ac3\u6536\u7d0d\u2728\u5c08\u4eba\u7ba1\u7406\u2728",
	// 				"type": 1,
	// 				"post_id": 14842358,
	// 				"price": "11,500",
	// 				"price_unit": "\u5143\/\u6708",
	// 				"photo_list": [
	// 					"https:\/\/img1.591.com.tw\/house\/2023\/04\/20\/168195553168053709.jpg!510x400.jpg",
	// 					"https:\/\/img2.591.com.tw\/house\/2023\/04\/20\/168195553107460806.jpg!510x400.jpg",
	// 					"https:\/\/img1.591.com.tw\/house\/2023\/04\/20\/168195553125013701.jpg!510x400.jpg",
	// 					"https:\/\/img2.591.com.tw\/house\/2023\/04\/20\/168195552998814402.jpg!510x400.jpg",
	// 					"https:\/\/img2.591.com.tw\/house\/2023\/04\/20\/168195553448700805.jpg!510x400.jpg"
	// 				],
	// 				"section_name": "\u897f\u5c6f\u5340",
	// 				"street_name": "\u5bcc\u5f37\u5df7",
	// 				"rent_tag": [
	// 					{
	// 						"id": 2,
	// 						"name": "\u8fd1\u6377\u904b"
	// 					}
	// 				],
	// 				"area": "10",
	// 				"surrounding": {
	// 					"type": "subway_station",
	// 					"desc": "\u8ddd\u6377\u904b\u6587\u83ef\u9ad8\u4e2d\u7ad9",
	// 					"distance": "425\u516c\u5c3a"
	// 				},
	// 				"community": "\u54c1\u5275\u5bcc\u5bcc1",
	// 				"room_str": "",
	// 				"is_video": 0,
	// 				"preferred": 1,
	// 				"kind": 2
	// 			}
	// 		]
	// 	}
	// }`)
	// data := house591.HouseStructure{}
	// json.Unmarshal([]byte(s), &data)
	// fmt.Printf("Status: %s \n", data.Data["topData"].([]interface{})[0].(map[string]interface{})["kind"])

	crawler, err := house591.Default()
	if err != nil {
		log.Fatalln("Initial 591 crawler failed:", err)
	}

	data, err := crawler.FetchHouses(house591.DefaultOptions())
	if err != nil {
		log.Fatalln("Fetch data failed:", err)
	}
	fmt.Println(data.Status)
	fmt.Println(data.Data.TopData)
}

func broadcast() {
	bot, err := linebotClient.Init()
	if err != nil {
		log.Fatalln("Initial linebot client error:", err)
	}

	bot.BroadcastMessage(linebot.NewTextMessage("Hi")).Do()
	bot.BroadcastMessage(
		linebot.NewImagemapMessage(
			"https://img1.591.com.tw/house/2023/08/17/169228532221864404.jpg!510x400.jpg",
			"Test",
			linebot.ImagemapBaseSize{Width: 1040, Height: 650},
			linebot.NewMessageImagemapAction("uri", "https://www.facebook.com/", linebot.ImagemapArea{X: 0, Y: 0, Width: 520, Height: 325}),
			linebot.NewMessageImagemapAction("uri", "https://www.google.com.tw/", linebot.ImagemapArea{X: 520, Y: 0, Width: 520, Height: 325}),
			linebot.NewMessageImagemapAction("message", "Facebook", linebot.ImagemapArea{X: 0, Y: 325, Width: 520, Height: 325}),
			linebot.NewMessageImagemapAction("message", "Google", linebot.ImagemapArea{X: 520, Y: 325, Width: 520, Height: 325}),
		),
	).Do()
}
