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

	c.OnHTML("div.structItem-cell--main", func(h *colly.HTMLElement) {
		// name
		fmt.Println(h.ChildText("div.structItem-title"))
		// author
		fmt.Println(h.ChildText("a[data-xf-init=member-tooltip]"))
		// url
		fmt.Println(h.ChildAttr("a[data-xf-init=preview-tooltip]", "href"))

		fmt.Println()
	})

	c.Visit("https://forums.rpgmakerweb.com/index.php?forums/rgss3-scripts-rmvx-ace.35/")
}
