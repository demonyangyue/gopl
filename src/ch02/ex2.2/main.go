package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			f, err := strconv.ParseFloat(arg, 64)

			if err != nil {
				log.Fatalf("#{arg} is not a valid float number")
			}
			fmt.Println(Celsius(f))
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {

			value := scan.Text()
			f, err := strconv.ParseFloat(value, 64)

			if err != nil {
				log.Fatalf("#{value} is not a valid float number")
			}
			fmt.Println(Celsius(f))
		}
	}
}
