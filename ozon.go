package main

import (
	"fmt"
	"net/url"
	"strings"
)

func ozonParse(data string) {
	data = urlEncode(data)

	fmt.Println(data)

}

func urlEncode(data string) string {
	data = url.QueryEscape(data)
	data = strings.ReplaceAll(data, "+", "%20")
	return data
}
