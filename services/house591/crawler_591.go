package house591

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

// https://rent.591.com.tw/home/search/rsList?is_format_data=1&is_new_list=1&type=1&region=8&section=104,101,100,105&searchtype=1&other=pet,newPost&recom_community=1
type Crawler struct {
	colly        *colly.Collector
	homePage     string
	searchingUrl string
	csrfToken    string
	cookies      string
}

func Default() (*Crawler, error) {
	return &Crawler{
		colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"),
		),
		"https://rent.591.com.tw",
		"https://rent.591.com.tw/home/search/rsList",
		"",
		"",
	}, nil
}

func NewCrawler(collector *colly.Collector, homePage string, searchingUrl string) (*Crawler, error) {
	return &Crawler{
		collector,
		homePage,
		searchingUrl,
		"",
		"",
	}, nil
}

func (c *Crawler) GetCsrfToken() string {
	return c.csrfToken
}

func (c *Crawler) FetchCsrfToken() (string, error) {
	c.colly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visit URL:", r.URL)
	})

	c.colly.OnHTML("meta[name='csrf-token']", func(e *colly.HTMLElement) {
		c.csrfToken = e.Attr("content")
	})

	c.colly.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "Error:", err)
	})

	c.colly.OnResponse(func(r *colly.Response) {
		for i, e := range r.Headers.Values("Set-Cookie") {
			fmt.Println(i, e)
		}
		c.cookies = strings.Join(r.Headers.Values("Set-Cookie"), ";")
	})

	c.colly.Visit(c.homePage)

	return c.csrfToken, nil
}

func (c *Crawler) FetchHouses(o *Options) (*HouseStructure, error) {
	searchingUrl := c.searchingUrl + "?" + o.ToQueryString()

	_, err := c.FetchCsrfToken()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", searchingUrl, nil)
	if err != nil {
		return nil, err
	}
	c.cookies = strings.Replace(c.cookies, "urlJumpIp=1", "urlJumpIp="+o.Region, 1)
	c.cookies = strings.Replace(c.cookies, "urlJumpIpByTxt=%E5%8F%B0%E5%8C%97%E5%B8%82", "urlJumpIpByTxt=", 1)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", c.cookies)
	req.Header.Add("X-CSRF-TOKEN", c.csrfToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := HouseStructure{}
	json.Unmarshal(bodyBytes, &data)

	// json.NewDecoder(res.Body).Decode(&data)
	return &data, nil
}
