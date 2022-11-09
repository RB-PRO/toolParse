package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func sitilinkParse(input string) {
	input = strings.Replace(input, " ", "+", -1)
	c := colly.NewCollector()

	c.OnHTML("div[class=ProductCardCategoryList__grid] div[class^=product_data__gtm-js]", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Attr("data-params"))
	})

	c.OnHTML("a[class=PaginationWidget__page js--PaginationWidget__page PaginationWidget__page_last PaginationWidget__page-link]", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Attr("data-page"))
	})

	c.Visit("https://www.citilink.ru/search/?text=" + input)
}
