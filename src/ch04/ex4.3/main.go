package main

import "fmt"

func main() {
	a := [5]int{1,2,4,3,5}
	reverse(&a)
	fmt.Printf("%v\n" , a)
}


//!+rev
// reverse reverses a slice of ints in place.
func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

