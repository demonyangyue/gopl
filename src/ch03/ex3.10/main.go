package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	length := len(s)

	for  i:= 0; i < length; i ++  {
		buf.WriteByte(s[i])
		if (length - 1 -i) %3 == 0 && (length - 1 -i) != 0 {
			buf.WriteByte(',')
		}
	}

	return buf.String()

}

//!-


