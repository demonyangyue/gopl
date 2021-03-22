package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := []byte("hello 中国")
	res := string(reverseUTF8(s))
	fmt.Println(res)
	
}

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

func reverseUTF8(s []byte) []byte {

	for i := 0; i < len(s);  {
		_, size := utf8.DecodeRune(s[i:])
		rev(s[i : i+size])
		i += size
	}
	rev(s)
	return s
}
