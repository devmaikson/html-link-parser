package parser

import (
	"os"

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

		if linkFound {
			linkItem.Text = token.Data
			linkList = append(linkList, linkItem)

			linkFound = false
		}

		if token.Data == "a" && tt == html.StartTagToken {
			for _, v := range token.Attr {
				linkItem.Href = v.Val
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
