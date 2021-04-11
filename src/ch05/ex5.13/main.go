package main

import (
	"fmt"
	"github.com/demonyangyue/gopl/src/ch05/links"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string, root string) []string, worklist []string, rootUrl string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, rootUrl)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string, root string) []string {
	saveToLocal(url, root)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func saveToLocal(rawurl string, root string) error {
	if !strings.HasPrefix(rawurl, root) {
		return nil
	}

	url, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("bad url: %s", err)
	}

	//create dir

	dir := url.Host
	var filename string
	if filepath.Ext(filename) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = url.Path
	}
	fmt.Println(filename)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	//create file

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	//getUrlContent

	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//write to file

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	// Check for delayed write errors, as mentioned at the end of section 5.8.
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.

	for _ , arg := range os.Args[1:] {
		breadthFirst(crawl, []string{arg}, arg)
	}
}

//!-main


