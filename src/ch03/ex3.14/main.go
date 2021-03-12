package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	s1 := os.Args[1]
	s2 := os.Args[2]

	fmt.Println(isAnagram(s1, s2))
	
}

func isAnagram(s1, s2 string) bool {

	s1Freq := make(map[rune]int)
	s2Freq := make(map[rune]int)

	for _, v1 := range s1 {
		s1Freq[v1] ++
	}

	for _, v2 := range s2 {
		s2Freq[v2] ++
	}

	return reflect.DeepEqual(s1Freq, s2Freq)

}

