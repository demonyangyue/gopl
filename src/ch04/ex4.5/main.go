package main

import "fmt"

func main() {

	ss := []string{"aa", "cc", "bb", "bb", "dd", "dd", "dd", "ee"}
	ss = dedup(ss)
	fmt.Println(ss)
	
}

func dedup(ss []string ) []string {

	//orig代表原始slice的下标，new代表新slice的下标
	new := 0
	for orig := 0 ; orig < len(ss); orig++ {
		if orig == 0 ||  ss[orig] != ss[orig-1] {
			ss[new] = ss[orig]
			new ++
		}
	}

	return ss[:new]
}
