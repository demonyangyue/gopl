package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(strings.NewReader("<html><script>hello</script><span>world</span></html>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "print text node content : %v\n", err)
		os.Exit(1)
	}
	for _, content := range visit(nil, doc) {
		fmt.Println(content)
	}
}

//!-main

//!+visit
func visit(contents []string, n *html.Node) []string {
	//跳过<script>和<style>节点
	if n.Data == "script" || n.Data == "style" {
		return contents
	}

	if n.Type == html.TextNode {
		contents = append(contents, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contents = visit(contents, c)
	}

	return contents
}

