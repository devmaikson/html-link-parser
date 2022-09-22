package main

import (
	"fmt"
	"html-link/parser"
)

func main() {
	f, err := parser.ReadHTMLDoc("resources/example1.html")

	if err != nil {
		panic(err)
	}

	ll := parser.ParseFile(f)
	fmt.Println()

	fmt.Println("Printing: Slice:", ll)
	for _, v := range ll {
		fmt.Println("Href: ", v.Href)
		fmt.Println("Text: ", v.Text)
	}
}
