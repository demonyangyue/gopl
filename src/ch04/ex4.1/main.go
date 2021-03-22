package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(countDiff(c1[:], c2[:]))
}

func countDiff(a, b []byte) int  {

	cnt := 0
	length := len(a)

	for i := 0; i < length; i++ {
		cnt += countPopulation(a[i] ^ b[i])
	}

	return cnt

}

func countPopulation(x byte) int {
	res := 0
	currentValue := x

	for ; currentValue!= 0; {
		currentValue = currentValue & (currentValue - 1)
		res += 1
	}

	return res
}

