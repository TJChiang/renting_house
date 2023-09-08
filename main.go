package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"renting_house/internal/line_message"
	"renting_house/services/house591"
	"strconv"

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
	bot, err := line_message.Init()
	if err != nil {
		log.Fatalln("Initial linebot client error:", err)
		return
	}

	var flex *int
	flex = new(int)
	*flex = 1
	k := 0
	var bubbleContainerList []*linebot.BubbleContainer
	for _, v := range data.Data.Data {
		bubbleContainerList = append(bubbleContainerList, &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Size: linebot.FlexBubbleSizeTypeKilo,
			Header: &linebot.BoxComponent{
				Type:       linebot.FlexComponentTypeBox,
				Layout:     linebot.FlexBoxLayoutTypeHorizontal,
				PaddingAll: "0px",
				Contents: []linebot.FlexComponent{
					&linebot.ImageComponent{
						URL:         v.ImageList[0],
						Size:        linebot.FlexImageSizeTypeFull,
						AspectMode:  linebot.FlexImageAspectModeTypeCover,
						AspectRatio: "150:196",
						Gravity:     linebot.FlexComponentGravityTypeCenter,
						Flex:        flex,
					},
					&linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Flex:   flex,
						Contents: []linebot.FlexComponent{
							&linebot.ImageComponent{
								URL:         v.ImageList[1],
								Size:        linebot.FlexImageSizeTypeFull,
								AspectMode:  linebot.FlexImageAspectModeTypeCover,
								AspectRatio: "150:98",
								Gravity:     linebot.FlexComponentGravityTypeCenter,
							},
							&linebot.ImageComponent{
								URL:         v.ImageList[1],
								Size:        linebot.FlexImageSizeTypeFull,
								AspectMode:  linebot.FlexImageAspectModeTypeCover,
								AspectRatio: "150:98",
								Gravity:     linebot.FlexComponentGravityTypeCenter,
							},
						},
					},
				},
			},
			Body: &linebot.BoxComponent{
				Type:          linebot.FlexComponentTypeBox,
				Layout:        linebot.FlexBoxLayoutTypeVertical,
				PaddingTop:    "15px",
				PaddingBottom: "10px",
				PaddingStart:  "20px",
				PaddingEnd:    "20px",
				Contents: []linebot.FlexComponent{
					&linebot.BoxComponent{
						Type:    linebot.FlexComponentTypeBox,
						Layout:  linebot.FlexBoxLayoutTypeVertical,
						Spacing: linebot.FlexComponentSpacingTypeSm,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Size:   linebot.FlexTextSizeTypeLg,
								Wrap:   true,
								Text:   v.Title,
								Color:  "#000000",
								Weight: linebot.FlexTextWeightTypeBold,
							},
						},
					},
					&linebot.BoxComponent{
						Type:    linebot.FlexComponentTypeBox,
						Layout:  linebot.FlexBoxLayoutTypeVertical,
						Spacing: linebot.FlexComponentSpacingTypeSm,
						Margin:  linebot.FlexComponentMarginTypeMd,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:  linebot.FlexComponentTypeText,
								Size:  linebot.FlexTextSizeTypeSm,
								Text:  v.KindName + " | " + v.Room + " | " + v.Area + "坪 | " + v.Floor,
								Color: "#090909",
							},
							&linebot.TextComponent{
								Type:  linebot.FlexComponentTypeText,
								Size:  linebot.FlexTextSizeTypeSm,
								Text:  v.SectionName + v.StreetName,
								Color: "#090909",
							},
						},
					},
					&linebot.BoxComponent{
						Type:    linebot.FlexComponentTypeBox,
						Layout:  linebot.FlexBoxLayoutTypeVertical,
						Spacing: linebot.FlexComponentSpacingTypeSm,
						Margin:  linebot.FlexComponentMarginTypeMd,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type:   linebot.FlexComponentTypeText,
								Size:   linebot.FlexTextSizeTypeLg,
								Weight: linebot.FlexTextWeightTypeBold,
								Text:   v.Price + " " + v.PriceUnit,
								Color:  "#E60012",
							},
						},
					},
					&linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Margin: linebot.FlexComponentMarginTypeMd,
						Contents: []linebot.FlexComponent{
							&linebot.ButtonComponent{
								Type:   linebot.FlexComponentTypeButton,
								Action: linebot.NewURIAction("開啟", "https://rent.591.com.tw/home/"+strconv.Itoa(v.PostId)),
							},
						},
					},
				},
			},
		})
		k++
		if k >= 12 {
			break
		}
	}

	contents := &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: bubbleContainerList,
	}

	res, err := bot.Client.BroadcastMessage(linebot.NewFlexMessage("carousel", contents)).Do()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(res)
}
