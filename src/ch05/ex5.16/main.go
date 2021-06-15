package main

import (
	"fmt"
	"strings"
)

func variadicJoin(sep string, vals ...string) string  {
	return strings.Join(vals, sep)

}

func main() {
	fmt.Println(variadicJoin(",", "hi", "world"))
	
}
