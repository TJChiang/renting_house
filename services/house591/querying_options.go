package house591

import (
	"net/url"
	"strings"
)

// ?is_format_data=1&is_new_list=1&type=1&region=8&section=104,101,100,105&searchtype=1&other=pet,newPost&recom_community=1
type Options struct {
	isFormatData   string
	isNewList      string
	Type           string
	Region         string
	Section        []string
	SearchType     string
	Other          []string
	RecomCommunity string
}

func NewOptions(
	Region string,
	Section []string,
	SearchType string,
	Other []string,
	RecomCommunity string,
) (*Options, error) {
	return &Options{
		isFormatData:   "1",
		isNewList:      "1",
		Type:           "1",
		Region:         Region,
		Section:        Section,
		SearchType:     SearchType,
		Other:          Other,
		RecomCommunity: RecomCommunity,
	}, nil
}

func DefaultOptions() *Options {
	return &Options{
		"1",
		"1",
		"1",
		"8",
		[]string{"104", "101", "100", "105"},
		"1",
		nil,
		"",
	}
}

func (o *Options) ToQueryString() string {
	v := url.Values{}
	v.Set("is_format_data", o.isFormatData)
	v.Set("is_new_list", o.isNewList)
	v.Set("type", o.Type)
	v.Set("region", o.Region)

	if len(o.Section) > 0 {
		v.Set("section", strings.Join(o.Section, ","))
	}
	if len(o.SearchType) > 0 {
		v.Set("searchtype", o.SearchType)
	}
	if len(o.Other) > 0 {
		v.Set("other", strings.Join(o.Other, ","))
	}
	if len(o.RecomCommunity) > 0 {
		v.Set("recom_community", o.RecomCommunity)
	}

	return v.Encode()
}
