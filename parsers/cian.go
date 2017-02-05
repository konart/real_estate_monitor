package parsers

import (
	"net/url"
	"net/http"
	"log"
	"golang.org/x/net/html"
	"io"
	"strings"
	"unicode"
)

func GetFlatPage(u *url.URL) io.ReadCloser {
	resp, err := http.Get(u.String())
	if err != nil {
		log.Printf("Could not fetch page %s", u.String())
	}

	page := resp.Body
	defer resp.Body.Close()
	return page
}

func ParseFlatPage(page io.Reader) map[string]string {
	height := findHeight(page)
	return map[string]string{"rooms":"1", "height":height}
}

func findHeight(page io.Reader) (height string) {
	var isHeightRow bool
	tokenizer := html.NewTokenizer(page)
	for {
		tt := tokenizer.Next()

		switch {
		case (tt == html.TextToken) && !isHeightRow:
			t := tokenizer.Token()
			isHeightRow = t.String() == "Высота потолков:"
			if isHeightRow {continue}
		case (tt == html.TextToken) && isHeightRow:
			tokenizer.Next()
			tokenizer.Next()
			height = strings.TrimSuffix(spaceMap(tokenizer.Token().String()), "м")
			return height
		default:
			continue
		}
	}

	return ""
}

func findNumberOfRooms(page io.Reader) (rooms string) {
	return "2"
}

func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}