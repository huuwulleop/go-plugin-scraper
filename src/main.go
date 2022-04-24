package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Item struct {
	Name   string `json:"name"`
	Author string `json:"price"`
	Url    string `json:"url"`
}

func main() {
	fmt.Println("Test")

	c := colly.NewCollector(
		colly.AllowedDomains("forums.rpgmakerweb.com"),
	)

	c.OnHTML("div[data-author=IMP1]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://forums.rpgmakerweb.com/index.php?forums/rgss3-scripts-rmvx-ace.35/")
}
