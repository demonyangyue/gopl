package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	wd , _ := os.Getwd()
	path, _ := filepath.Abs(filepath.Join(wd, "/src/ch04/ex4.9/main.go"))
	wordFreq(path)
}

func wordFreq(path string)  {
	freq := make(map[string]int)

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("can't find file %s", path)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		freq[word]++
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, scanner.Err())
		os.Exit(1)
	}

	for word, n := range freq {
		fmt.Printf("%-30s %d\n", word, n)
	}

}
