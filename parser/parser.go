package parser

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

type LinkList []Link

type Link struct {
	Href string
	Text string
}

func ParseFile(r *os.File) LinkList {
	var linkItem Link
	linkList := make(LinkList, 0)
	linkFound := false
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		token := z.Token()

		if tt == html.EndTagToken && linkFound {
			linkFound = false
			linkList = append(linkList, linkItem)
			linkItem.Text = ""
		}

		if linkFound && tt == html.TextToken {
			if linkItem.Text == "" {
				linkItem.Text = strings.TrimSpace(token.Data)
			} else {
				linkItem.Text = linkItem.Text + " " + strings.TrimSpace(token.Data)
			}
		}

		if token.Data == "a" && tt == html.StartTagToken {
			for _, v := range token.Attr {
				if v.Key == "href" {
					linkItem.Href = v.Val
				}
			}
			linkFound = true
		}

	}
	return linkList
}

func ReadHTMLDoc(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}
