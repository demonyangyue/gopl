package main

import (
	"fmt"
)

func max(vals ...int)(int, error)  {
	result := 0
	if len(vals) == 0 {
		return result, fmt.Errorf("invaild argment length")
	}
	for _, val := range vals {
		if val > result {
			result = val
		}
	}

	return result, nil


}

func min(vals ...int)(int, error)  {
	result := 0
	if len(vals) == 0 {
		return result, fmt.Errorf("invaild argment length")
	}
	for _, val := range vals {
		if val < result {
			result = val
		}
	}

	return result, nil


}
func main() {
	fmt.Println(max(1,3,5,7))
	fmt.Println(min(1,3,5,7))

}
