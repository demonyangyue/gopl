package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: ex5.8 HTML_FILE ID")
	}
	filename := os.Args[1]
	id := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", ElementByID(doc, id))

}

func ElementByID(doc *html.Node, id string) *html.Node {
	return  forEachNode(doc, id, startElement, endElement)
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string,  pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil && !pre(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, id, pre, post)
	}

	if post != nil && !post(n, id){
		return n
	}

	return nil
}

//!-forEachNode


func startElement(n *html.Node, id string) bool {

	if len(n.Attr) > 0 {
		for _, attribute := range n.Attr {
			if attribute.Key == "id" && attribute.Val == id {
				return false
			}
		}
	}

	return true

}


func endElement(n *html.Node, id string) bool {
	return true
}

//!-startend


