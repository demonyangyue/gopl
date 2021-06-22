package main

import (
	"fmt"
	"github.com/demonyangyue/gopl/src/ch07/eval"
)

func main() {
	exp, _ := eval.Parse("sqrt(2)")
	fmt.Println(exp)

	exp, _ = eval.Parse("sqrt(A / pi)")
	fmt.Println(exp)

	exp, _ = eval.Parse("pow(x, 3) + pow(y, 3)")
	fmt.Println(exp)

	exp, _ = eval.Parse("5 / 9 * (F - 32)")
	fmt.Println(exp)
}
