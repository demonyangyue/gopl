// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"github.com/demonyangyue/gopl/src/ch05/links"
	"log"
	"os"

)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

var depth = flag.Int("depth", 3, "max depth")

type  WorkItem struct {
	items []string
	current_depth int
}

type UnseenLink struct {
	link string
	current_depth int
}
//!+
func main() {
	flag.Parse()

	worklist := make(chan *WorkItem)  // lists of URLs, may have duplicates
	unseenLinks := make(chan *UnseenLink) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- &WorkItem{os.Args[1:], 0}}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for unseenLink := range unseenLinks {
				if unseenLink.current_depth >= *depth {
					continue
				}
				foundLinks := crawl(unseenLink.link)
				go func() { worklist <- &WorkItem{foundLinks, unseenLink.current_depth + 1} }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list.items {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- &UnseenLink{link, list.current_depth}
			}
		}
	}
}

//!-
