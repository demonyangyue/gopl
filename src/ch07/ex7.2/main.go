package main

import (
	"bytes"
	"fmt"
	"io"
)

type ByteCounter struct {
	w io.Writer
	val int64
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	len, err := c.w.Write(p)
	c.val += int64(len)
	return len, err
}

func CountingWrite(w io.Writer) (io.Writer, *int64)  {
	c := &ByteCounter{w, 0}
	return c, &c.val
}

func main() {
	c, cntp := CountingWrite(&bytes.Buffer{})
	fmt.Fprintf(c, "hello world")
	fmt.Println(*cntp)
}
