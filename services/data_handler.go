package services

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"os"
	"renting_house/services/house591"
	"strconv"
)

func GetBubbleContainer(data house591.MainDataElement) *linebot.BubbleContainer {
	return &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Size: linebot.FlexBubbleSizeTypeKilo,
		Header: &linebot.BoxComponent{
			Type:       linebot.FlexComponentTypeBox,
			Layout:     linebot.FlexBoxLayoutTypeHorizontal,
			PaddingAll: "0px",
			Contents:   convertMainDataToFlexComponent(data),
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
							Text:   data.Title,
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
							Text:  data.KindName + " | " + data.Room + " | " + data.Area + "坪 | " + data.Floor,
							Color: "#090909",
						},
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Size:  linebot.FlexTextSizeTypeSm,
							Text:  data.SectionName + data.StreetName,
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
							Text:   data.Price + " " + data.PriceUnit,
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
							Action: linebot.NewURIAction("開啟", os.Getenv("RENT_HOUSE_ITEM_URL")+strconv.Itoa(data.PostId)),
						},
					},
				},
			},
		},
	}
}

func convertMainDataToFlexComponent(data house591.MainDataElement) []linebot.FlexComponent {
	listLength := len(data.ImageList)

	switch {
	case listLength > 2:
		return []linebot.FlexComponent{
			&linebot.ImageComponent{
				URL:         data.ImageList[0],
				Size:        linebot.FlexImageSizeTypeFull,
				AspectMode:  linebot.FlexImageAspectModeTypeCover,
				AspectRatio: "150:196",
				Gravity:     linebot.FlexComponentGravityTypeCenter,
			},
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.ImageComponent{
						URL:         data.ImageList[1],
						Size:        linebot.FlexImageSizeTypeFull,
						AspectMode:  linebot.FlexImageAspectModeTypeCover,
						AspectRatio: "150:98",
						Gravity:     linebot.FlexComponentGravityTypeCenter,
					},
					&linebot.ImageComponent{
						URL:         data.ImageList[2],
						Size:        linebot.FlexImageSizeTypeFull,
						AspectMode:  linebot.FlexImageAspectModeTypeCover,
						AspectRatio: "150:98",
						Gravity:     linebot.FlexComponentGravityTypeCenter,
					},
				},
			},
		}
	case listLength > 1:
		return []linebot.FlexComponent{
			&linebot.ImageComponent{
				URL:         data.ImageList[0],
				Size:        linebot.FlexImageSizeTypeFull,
				AspectMode:  linebot.FlexImageAspectModeTypeCover,
				AspectRatio: "150:196",
				Gravity:     linebot.FlexComponentGravityTypeCenter,
			},
			&linebot.ImageComponent{
				URL:         data.ImageList[1],
				Size:        linebot.FlexImageSizeTypeFull,
				AspectMode:  linebot.FlexImageAspectModeTypeCover,
				AspectRatio: "150:196",
				Gravity:     linebot.FlexComponentGravityTypeCenter,
			},
		}
	case listLength > 0:
		return []linebot.FlexComponent{
			&linebot.ImageComponent{
				URL:         data.ImageList[0],
				Size:        linebot.FlexImageSizeTypeFull,
				AspectMode:  linebot.FlexImageAspectModeTypeCover,
				AspectRatio: "300:196",
				Gravity:     linebot.FlexComponentGravityTypeCenter,
			},
		}
	}

	return []linebot.FlexComponent{
		&linebot.ImageComponent{
			URL:         data.ImageList[0],
			Size:        linebot.FlexImageSizeTypeFull,
			AspectMode:  linebot.FlexImageAspectModeTypeCover,
			AspectRatio: "300:196",
			Gravity:     linebot.FlexComponentGravityTypeCenter,
		},
	}
}
