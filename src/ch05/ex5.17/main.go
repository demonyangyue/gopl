package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func elementsByTagName(node *html.Node, tags ...string) []*html.Node {
	if node == nil || len(tags) == 0 {
		return nil
	}

	var result []*html.Node

	for _, tag := range tags {

		if node.Type == html.ElementNode && node.Data == tag {
			result =  append(result, node)
		}
	}


	for c := node.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, elementsByTagName(c, tags...)...)
	}

	return result
	
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementsByTagName: %v\n", err)
		os.Exit(1)
	}

	var result  []*html.Node;

	result = elementsByTagName(doc, "h1", "h2" )

	for _, node := range result {
		fmt.Printf("%s \n", node.FirstChild.Data)
	}
}

