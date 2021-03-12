package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		sign, intPart, fracPart := splitNumber(os.Args[i])
		fmt.Printf("  %s\n", sign + commaInt(intPart) + fracPart)
	}
}

//!+
// comma_int inserts commas in a non-negative decimal integer string.
func commaInt(s string) string {
	var buf bytes.Buffer
	length := len(s)

	for  i:= 0; i < length; i ++  {
		buf.WriteByte(s[i])
		if (length - 1 -i) %3 == 0 && (length - 1 -i) != 0 {
			buf.WriteByte(',')
		}
	}

	return buf.String()

}

//将数字切分成符号、整数、小数三部分
func splitNumber(s string) (string, string, string) {

	sign, intPart, fracPart := "", "", ""

	if strings.Index(s, "+") == 0 || strings.Index(s, "-") == 0 {
		sign = string(s[0])
		s = s[1:]
	}

	fracIndex := strings.Index(s, ".")
	if fracIndex > 0 {
		fracPart = s[fracIndex:]
		s = s[:fracIndex]
	}

	intPart = s

	return sign, intPart, fracPart

}

//提取小数部分
//!-


