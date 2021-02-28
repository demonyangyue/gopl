package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = addPrefixIfMissing(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b , err := ioutil.ReadAll(resp.Body)
		status := resp.StatusCode
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n",  b)
		fmt.Printf("%d\n",  status)
	}
}

func addPrefixIfMissing(url string) string {

	if  strings.HasPrefix(url, "http://") {
		return url
	}
	return "http://" + url
}

