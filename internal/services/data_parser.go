package services

import (
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"strings"
)

type Data struct {
	Title  string
	Artist string
}

func parseAppleMusic(url string) (string, string) {

	var data Data

	c := colly.NewCollector(
		colly.AllowedDomains("music.apple.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"))

	c.OnHTML("meta[name='apple:title']", func(e *colly.HTMLElement) {
		data.Title = e.Attr("content")
	})

	c.OnHTML("span.multiline-clamp__text a[Data-testid='click-action']", func(e *colly.HTMLElement) {
		data.Artist = strings.TrimSpace(e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Ошибка запроса:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Println("Ошибка при посещении страницы", err)
	}

	return data.Title, data.Artist
}

func parseVkMusic(url string) (string, string) {

	var data Data

	c := colly.NewCollector(
		colly.AllowedDomains("vk.com", "login.vk.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	c.OnHTML("div.AudioPlaylistSnippet__title", func(e *colly.HTMLElement) {
		data.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("div.AudioPlaylistSnippet__author", func(e *colly.HTMLElement) {
		data.Artist = strings.TrimSpace(e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Ошибка запроса:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Println("Ошибка при посещении страницы:", err)
	}

	return data.Title, data.Artist
}

func parseSpotify(url string) (string, string) {

	var data Data

	c := colly.NewCollector(
		colly.AllowedDomains("open.spotify.com"))

	c.OnHTML("meta[property='og:title']", func(e *colly.HTMLElement) {
		data.Title = e.Attr("content")
	})

	c.OnHTML("meta[name='music:musician_description']", func(e *colly.HTMLElement) {
		data.Artist = e.Attr("content")
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Ошибка запроса:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Println("Ошибка при посещении страницы:", err)
	}

	return data.Title, data.Artist

}

func parseYandexMusic(url string) (string, string) {

	var data Data

	c := colly.NewCollector(
		colly.AllowedDomains("music.yandex.ru"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	c.OnHTML("div.sidebar__title a", func(e *colly.HTMLElement) {
		data.Title = strings.TrimSpace(e.Text)
	})

	c.OnHTML("div.page-album__artists-short a", func(e *colly.HTMLElement) {
		data.Artist = strings.TrimSpace(e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Ошибка запроса:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Println("Ошибка при посещении страницы:", err)
	}

	return data.Title, data.Artist
}

func GetData(rawURL string) (string, string) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Println("Ошибка парсинга URL:", err)
	}

	var title, artist string

	host := parsedURL.Host
	path := parsedURL.Path

	switch {
	case strings.Contains(host, "vk.com") && strings.Contains(path, "/audio"):
		title, artist = parseVkMusic(rawURL)
	case strings.Contains(host, "music.apple.com"):
		title, artist = parseAppleMusic(rawURL)
	case strings.Contains(host, "music.yandex.ru"):
		title, artist = parseYandexMusic(rawURL)
	case strings.Contains(host, "spotify.com"):
		title, artist = parseSpotify(rawURL)
	default:
		log.Println("Неподдерживаемый сервис")
	}
	return title, artist
}
