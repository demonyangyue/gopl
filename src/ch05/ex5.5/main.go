package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return visit(words, images, n)
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}

func visit(words, images int, n *html.Node) (int, int) {

	switch n.Type {
	case html.TextNode:
		words += wordCount(n.Data)
	case html.ElementNode:
		if n.Data == "img" {
			images++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		words, images = visit(words, images, c)
	}
	return words, images
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: PROG URL")
	}
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("words: %d\nimages: %d\n", words, images)
}


