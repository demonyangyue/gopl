package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {

	for i, j:= 0, s.Len() - 1; i < j; i, j = i+1, j-1{
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}

	}
	return true
}

func main() {

	s := []string{"a", "c", "b", "c","a"}
	fmt.Println(IsPalindrome(sort.StringSlice(s)))
}
