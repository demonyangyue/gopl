package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//每次除以256，统计余数的population count
func PopCountTable(x uint64) int {
	var sum, divRemain int
	divResult := int(x)
	for {
		if divResult == 0 {
			break
		}
		divRemain = divResult % 256
		divResult = divResult / 256
		sum += int(pc[divRemain])
	}
	return sum
}

func main() {
	fmt.Printf("%v\n", PopCountTable(256) )
	fmt.Printf("%v\n", PopCountTable(1000) )
}
