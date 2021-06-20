package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func NewReader(s string) *Reader { return &Reader{s, 0, -1} }

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func main() {
	doc, err := html.Parse(NewReader("<h1>hello</h1>"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "countElements: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(doc.FirstChild.Data)


}
