package services

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
)

func parseVkMusicLink(title, artist string) string {

	var link string

	searchQuery := fmt.Sprintf("%s %s", title, artist)
	encodedQuery := url.QueryEscape(searchQuery)
	searchURL := fmt.Sprintf("https://vk.com/search?q=%s", encodedQuery)

	c := colly.NewCollector(
		colly.AllowedDomains("vk.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (HTML, like Gecko) Chrome/111.0.0.0 Safari/537.3"))

	c.OnHTML("a.audio_row__cover", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.Contains(href, "audio") {
			link = "https://vk.com" + href
			e.Request.Abort()
		}
	})

	err := c.Visit(searchURL)
	if err != nil {
		fmt.Println("Ошибка при посещении страницы VK", err)
		return ""
	}
	return link
}
