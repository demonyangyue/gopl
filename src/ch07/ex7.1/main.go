package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//!+bytecounter

type ByteCounter struct {
	val int
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	c.val += ByteCounter{len(p)}.val // convert int to ByteCounter
	return len(p), nil
}

type WordCounter struct {
	val int
}

func (w *WordCounter) Write(p []byte) (int, error)  {
	w.val += retCount(p, bufio.ScanWords)
	return len(p), nil
}

type LineCounter struct {
	val int
}

func (l *LineCounter) Write(p []byte) (int, error)  {
	l.val += retCount(p, bufio.ScanLines)
	return len(p), nil
}


func retCount(p []byte, fn bufio.SplitFunc) (count int) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(fn)
	count = 0
	for scanner.Scan() {
		count ++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input ", err)
	}
	return
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = ByteCounter{0} // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main

	var w WordCounter
	fmt.Fprintf(&w, "hello world")
	fmt.Println(w)

	var l LineCounter
	fmt.Fprintln(&l, "hello world \n nice boy\n")
	fmt.Println(l)
}

