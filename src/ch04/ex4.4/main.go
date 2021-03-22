package main

import (
	"fmt"
	"log"
)

func main() {
	s := []int{1,2,3,4,5}
	fmt.Println(rotate(s, 2))

}

func rotate(s []int ,  n int) []int {
	if s == nil || n < 0 {
		log.Fatal("invalid param")
	}
	return append(s[n:] , s[:n]... )
}