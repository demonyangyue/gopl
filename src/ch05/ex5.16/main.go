package main

import (
	"fmt"
	"strings"
)

func variadic_join(sep string, vals ...string) string  {
	return strings.Join(vals, sep)

}

func main() {
	fmt.Println(variadic_join(",", "hi", "world"))
	
}
