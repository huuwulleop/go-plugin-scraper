package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Item struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Url    string `json:"url"`
}

func main() {
	fmt.Println("Test")

	c := colly.NewCollector(
		colly.AllowedDomains("forums.rpgmakerweb.com"),
	)

	items := make([]Item, 0)

	c.OnHTML("div.structItem-cell--main", func(h *colly.HTMLElement) {
		// name
		fmt.Println(h.ChildText("div.structItem-title"))
		// author
		fmt.Println(h.ChildText("a[data-xf-init=member-tooltip]"))
		// url
		fmt.Println(h.ChildAttr("a[data-xf-init=preview-tooltip]", "href"))

		fmt.Println()

		item := Item{
			Name:   h.ChildText("div.structItem-title"),
			Author: h.ChildText("a[data-xf-init=member-tooltip]"),
			Url:    h.ChildAttr("a[data-xf-init=preview-tooltip]", "href"),
		}

		items = append(items, item)
	})

	c.Visit("https://forums.rpgmakerweb.com/index.php?forums/rgss3-scripts-rmvx-ace.35/")

	fmt.Println(items)

	content, err := json.Marshal(items)

	if err != nil {
		panic("Can't format to JSON!")
	}

	os.WriteFile("plugins.json", content, 0644)
}
