package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	width := flag.Int("w", 256, "hash width, 256 or 512")
	flag.Parse()

	var function func (b []byte) []byte

	switch *width {
	case 256:
		function = func(b []byte) []byte {
			h := sha256.Sum256(b)
			return h[:]
		}
	case 512:
		function = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("Unexpected width specified.")
	}

	b, error := ioutil.ReadAll(os.Stdin)
	if (error != nil) {
		log.Fatal("unexpected input")
	}

	fmt.Printf("%x", function(b))
}
