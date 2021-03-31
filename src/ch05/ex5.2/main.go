package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countElements: %v\n", err)
		os.Exit(1)
	}

	elementsCount := make(map[string]int)

	visit(elementsCount, doc)

	for name, count := range elementsCount {
		fmt.Printf("%s has %d elements\n", name, count)
	}
}

//!-main

//!+visit
func visit(elementsCount map[string]int, n *html.Node) {
	if n.Type == html.ElementNode  {
		elementsCount[n.Data] += 1
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(elementsCount, c)
	}
}

