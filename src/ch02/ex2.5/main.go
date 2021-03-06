package main

import "fmt"

//每次除以256，统计余数的population count
func countPopulation(x int) int {
	res := 0
	currentValue := x

	for ; currentValue!= 0; {
		currentValue = currentValue & (currentValue - 1)
		res += 1
	}

	return res;
}

func main() {
	fmt.Printf("%v\n", countPopulation(0) )
	fmt.Printf("%v\n", countPopulation(1) )
	fmt.Printf("%v\n", countPopulation(256) )
	fmt.Printf("%v\n", countPopulation(1000) )
}
