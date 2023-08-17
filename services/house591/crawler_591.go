package house591

import (
	"fmt"
	"log"
	"net/url"
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
		c.cookies = strings.Join(r.Headers.Values("Set-Cookie"), ";")
	})

	c.colly.Visit(c.homePage)

	return c.csrfToken, nil
}

func (c *Crawler) FetchHouses(query string) error {
	searchingUrl, err := url.Parse(c.searchingUrl)
	if err != nil {
		log.Fatalln("Invalid searching url:", err)
		return err
	}

}
