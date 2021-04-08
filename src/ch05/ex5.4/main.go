package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	//style sheet links
	if n.Type == html.ElementNode && n.Data == "link" {
		if checkAttrVal(n.Attr, "type", "text/css") {
			links = append(links, extractAttrVal(n.Attr, "href"))
		}
	}

	//image and script links

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func extractAttrVal(attributes []html.Attribute, key string) string {
	for _, attr := range attributes {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func checkAttrVal(attributes []html.Attribute, key string, val string) bool {
	for _, attr := range attributes {
		if attr.Key == key && attr.Val == val {
			return true
		}
	}
	return false

}

